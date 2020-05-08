package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	"github.com/xiaobudongzhang/micro-payment-srv/model/payment"
	proto "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

var (
	paymentService payment.Service
)

type Service struct {
}

func Init() {
	paymentService, _ = payment.GetService()
}

func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder]收到支付请求")

	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
