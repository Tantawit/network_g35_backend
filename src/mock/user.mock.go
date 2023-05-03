package mocks

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UserClientMock struct {
	mock.Mock
}

func (c *UserClientMock) FindOneLocalUser(ctx context.Context, in *pb.FindOneLocalUserRequest, opts ...grpc.CallOption) (*pb.FindOneLocalUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *UserClientMock) FindOneUserStudent(ctx context.Context, in *pb.FindOneUserStudentRequest, opts ...grpc.CallOption) (*pb.FindOneUserStudentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *UserClientMock) FindOneUserCUFreelance(_ context.Context, req *pb.FindOneUserCUFreelanceRequest, opts ...grpc.CallOption) (res *pb.FindOneUserCUFreelanceResponse, err error) {
	args := c.Called(req)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.FindOneUserCUFreelanceResponse)
	}

	return res, args.Error(1)
}

func (c *UserClientMock) Update(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *UserClientMock) FindUsersFromList(ctx context.Context, in *pb.FindUsersFromListRequest, opts ...grpc.CallOption) (*pb.FindUsersFromListResponse, error) {
	//TODO implement me
	panic("implement me")
}
