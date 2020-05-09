package payment

import (
	"fmt"
	"sync"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/xiaobudongzhang/micro-basic/common"
	invS "github.com/xiaobudongzhang/micro-inventory-srv/proto/inventory"
	ordS "github.com/xiaobudongzhang/micro-order-srv/proto/order"
)

var (
	s            *service
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)

type service struct {
}

type Service interface {
	PayOrder(orderId int64) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("getservice 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	invClient = invS.NewInventoryService("mu.micro.book.service.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("mu.micro.book.service.order", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
