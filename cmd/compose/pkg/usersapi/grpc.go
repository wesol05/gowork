package usersapi

import (
	"context"
	"fmt"

	pb "github.com/mikenai/gowork/internal/proto"

	"google.golang.org/grpc"
)

type ClientGRPC struct {
	grpc pb.UsersServiceClient
}

func NewClientGRPC(conn grpc.ClientConnInterface) *ClientGRPC {
	client := pb.NewUsersServiceClient(conn)
	return &ClientGRPC{
		grpc: client,
	}
}

func (cl *ClientGRPC) GetUser(ctx context.Context, id string) (User, error) {
	usr, err := cl.grpc.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return User{}, fmt.Errorf("failed to create request: %w", err)
	}

	return User{ID: usr.Id, Name: usr.Name}, nil
}
