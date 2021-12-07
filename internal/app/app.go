package app

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"prj/internal/config"
	"prj/internal/database"
	"prj/internal/entity"
	"prj/internal/service/auth"
	"prj/internal/service/order"
	authpb "prj/proto/auth"
	orderpb "prj/proto/order"
)

func Run() {
	cfg := config.GetConfig()
	db, err := database.CreateClient(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf(
			"%s:%s",
			cfg.Server.Host,
			cfg.Server.Port,
		),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.AutoMigrate(&entity.Order{}, &entity.OrderStatus{})
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{})
	orderpb.RegisterOrderServiceServer(s, &order.Service{DB: db})
	fmt.Println("Server is listening...")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
