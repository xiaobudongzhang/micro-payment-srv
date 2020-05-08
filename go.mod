module github.com/xiaobudongzhang/micro-payment-srv

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /data/ndemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /data/ndemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /data/ndemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /data/ndemo/micro-order-srv

require (
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200430044506-2451e30bf530
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	github.com/xiaobudongzhang/micro-basic v1.1.5
	github.com/xiaobudongzhang/micro-inventory-srv v1.0.0
	github.com/xiaobudongzhang/micro-order-srv v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.22.0
)
