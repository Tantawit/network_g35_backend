package service

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/service"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
)

func NewAuthService(client pb.AuthServiceClient) AuthService {
	logger := gosdk.NewLogger("auth-service")

	return &service.AuthServiceImpl{
		Client: client,
		Logger: logger,
	}
}

type AuthService interface {
	Validate(ctx context.Context, token string) (*dto.TokenPayloadAuth, error)
}
