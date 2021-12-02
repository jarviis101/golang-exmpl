package auth

import (
	"context"
	pb "prj/proto/auth"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
}

type Payload struct {
	Phone string `json:"phone"`
}

func (s *Service) Get(ctx context.Context, in *pb.Payload) (*pb.Token, error) {
	return &pb.Token{
		Token: "Random str",
	}, nil
}
