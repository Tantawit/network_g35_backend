package service

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/service"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
)

func NewUserService(client pb.UserServiceClient) UserService {
	logger := gosdk.NewLogger("user-service")

	return &service.UserServiceImpl{
		Logger: logger,
		Client: client,
	}
}

type UserService interface {
	FindOneCUFreelance(ctx context.Context, id string) (*pb.UserCUFreelance, error)
}
