package app

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"prj/internal/config"
	"prj/internal/database"
	"prj/internal/service/auth"
	pb "prj/proto/auth"
)

func Run() {
	cfg := config.GetConfig()
	_, err := database.CreateClient(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%s",
		cfg.Server.Host,
		cfg.Server.Port,
	),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &auth.Service{})
	fmt.Println("Server is listening...")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
