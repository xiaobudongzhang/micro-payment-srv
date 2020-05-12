package payment

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/common"
	"github.com/xiaobudongzhang/micro-plugins/db"

	invS "github.com/xiaobudongzhang/micro-inventory-srv/proto/inventory"
	ordS "github.com/xiaobudongzhang/micro-order-srv/proto/order"
)

func (s *service) PayOrder(orderId int64) (err error) {
	orderRsp, err := ordSClient.GetOrder(context.TODO(), &ordS.Request{
		OrderId: orderId,
	})

	if err != nil {
		log.Logf("[payorder]查询订单信息失败 orderid:%d, %s", orderId, err)
		return
	}

	if orderRsp == nil || !orderRsp.Success || orderRsp.Order == nil {
		err = fmt.Errorf("[payorder] 支付单不存在")
		log.Logf("[payorder]查询订单信息失败 orderId:%d,%s", orderId, err)
		return
	}

	if orderRsp.Order.State == common.InventoryHistoryStateOut {
		err = fmt.Errorf("[payorder] 订单已支付")
		log.Logf("[payorder]查询 订单已支付, orderId:%d", orderId)
		return
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Logf("[payorder] 事务开启失败", err.Error())
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	insertSQL := `insert into payment(user_id,book_id,order_id, inv_his_id,state) value (?, ?,?,?,?)`

	_, err = tx.Exec(insertSQL, orderRsp.Order.UserId, orderRsp.Order.BookId, orderRsp.Order.Id, orderRsp.Order.InvHistoryId, common.InventoryHistoryStateOut)
	if err != nil {
		log.Logf("[new]新增订单失败， %v, err:%s", orderRsp.Order, err)
		return
	}

	invRsp, err := invClient.Confirm(context.TODO(), &invS.Request{
		HistoryId: orderRsp.Order.InvHistoryId, HistoryState: 2,
	})

	if err != nil || invRsp == nil || !invRsp.Success {
		err = fmt.Errorf("[payorder] 确认入库失败 %s", err)
		log.Logf("%s", err)
		return
	}

	s.sendPayDoneEvt(orderId, common.InventoryHistoryStateOut)

	tx.Commit()
	return

}
