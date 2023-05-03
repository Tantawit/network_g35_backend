package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
)

type UserServiceImpl struct {
	Logger gosdk.Logger
	Client pb.UserServiceClient
}

func (s *UserServiceImpl) FindOneCUFreelance(ctx context.Context, id string) (*pb.UserCUFreelance, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.Client.FindOneUserCUFreelance(ctx, &pb.FindOneUserCUFreelanceRequest{Id: id})
	if err != nil {
		if status.Code(err) != codes.NotFound {
			s.Logger.Error(err).
				Keyword("user_id", id).
				Msg(err.Error())
		}
		return nil, err
	}

	return res.User, nil
}
