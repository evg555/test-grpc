package user

import (
	"context"
	"fmt"
	pb "grpc-server/pkg/api"
)

type GRPCServer struct {
	pb.UnimplementedUserServer
}

func (srv *GRPCServer) Create(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Result: fmt.Sprintf("User with name %s and age %d created", request.GetName(), request.GetAge()),
	}, nil
}
