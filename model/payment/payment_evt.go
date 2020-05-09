package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/util/log"
	proto "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

func (s *service) sendPayDoneEvt(orderId int64, state int32) {

	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendpaydoneevt] 发送支付事件 %+v\n", ev)

	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Logf("[sendpaydoneevt] 异常：%v", err)
	}
}
