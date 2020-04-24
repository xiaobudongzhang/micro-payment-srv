package main

import (
	"github.com/xiaobudongzhang/micro-payment-srv/handler"
	"github.com/xiaobudongzhang/micro-payment-srv/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	payment "github.com/xiaobudongzhang/micro-payment-srv/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.service.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.service.payment", service.Server(), new(subscriber.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
