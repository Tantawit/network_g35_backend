package service

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"time"
)

type AuthServiceImpl struct {
	Client pb.AuthServiceClient
	Logger gosdk.Logger
}

func (s *AuthServiceImpl) Validate(ctx context.Context, token string) (*dto.TokenPayloadAuth, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.Client.Validate(ctx, &pb.ValidateRequest{Token: token})
	if err != nil {
		return nil, err
	}

	return &dto.TokenPayloadAuth{
		UserId:   res.UserId,
		UserType: constant.UserType(res.UserType),
	}, nil
}
