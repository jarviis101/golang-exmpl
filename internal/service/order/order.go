package order

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"prj/internal/factory"
	pb "prj/proto/order"
)

type Service struct {
	pb.UnimplementedOrderServiceServer
	DB *gorm.DB
}

func (s *Service) CreateOrder(ctx context.Context, in *pb.OrderPayload) (*pb.OrderResponse, error) {
	fc := factory.CreateOrderFactory(s.DB)
	order, err := fc.Create(in)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &pb.OrderResponse{
		HashId: order.Name,
	}, nil
}

func (s *Service) PayOrder(ctx context.Context, in *pb.PaymentPayload) (*pb.PaymentResponse, error) {
	// TODO: paid order
	return &pb.PaymentResponse{
		Status: "paid",
	}, nil
}
