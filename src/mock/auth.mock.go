package mocks

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ContextMock struct {
	mock.Mock
	V               interface{}
	Status          int
	Header          map[string]string
	VerifyTicketDto *dto.VerifyTicket
	RefreshTokenDto *dto.RedeemNewToken
}

func (c *ContextMock) Query(key string) string {
	args := c.Called(key)

	return args.String(0)
}

func (c *ContextMock) SendStatus(status int) error {
	args := c.Called(status)

	return args.Error(0)
}

func (c *ContextMock) Next() error {
	args := c.Called()

	return args.Error(0)
}

func (c *ContextMock) UserContext() context.Context {
	return context.Background()
}

func (c *ContextMock) Context() context.Context {
	return context.Background()
}

func (c *ContextMock) UserType() constant.UserType {
	args := c.Called()

	return args.Get(0).(constant.UserType)
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	switch v.(type) {
	case *dto.VerifyTicket:
		*v.(*dto.VerifyTicket) = *c.VerifyTicketDto
	case *dto.RedeemNewToken:
		*v.(*dto.RedeemNewToken) = *c.RefreshTokenDto
	case *dto.RegisterDto:
		*v.(*dto.RegisterDto) = *args.Get(0).(*dto.RegisterDto)
	case *dto.LoginDto:
		*v.(*dto.LoginDto) = *args.Get(0).(*dto.LoginDto)
	case *dto.ChangePasswordDto:
		*v.(*dto.ChangePasswordDto) = *args.Get(0).(*dto.ChangePasswordDto)
	}

	return args.Error(1)
}

func (c *ContextMock) JSON(status int, v interface{}) {
	c.V = v
	c.Status = status
}

func (c *ContextMock) UserID() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Token() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) StoreValue(key string, val string) {
	_ = c.Called(key, val)

	c.Header[key] = val
}

func (c *ContextMock) Method() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Path() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) AuthInfo() *dto.AuthInfo {
	args := c.Called()

	return args.Get(0).(*dto.AuthInfo)
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) ChangePassword(_ context.Context, in *pb.ChangePasswordRequest, _ ...grpc.CallOption) (res *pb.ChangePasswordResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.ChangePasswordResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Logout(_ context.Context, in *pb.LogoutRequest, _ ...grpc.CallOption) (res *pb.LogoutResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.LogoutResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Register(_ context.Context, in *pb.RegisterRequest, _ ...grpc.CallOption) (res *pb.RegisterResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.RegisterResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Login(_ context.Context, in *pb.LoginRequest, _ ...grpc.CallOption) (res *pb.LoginResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.LoginResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) VerifyTicket(_ context.Context, in *pb.VerifyTicketRequest, _ ...grpc.CallOption) (res *pb.VerifyTicketResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.VerifyTicketResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Validate(_ context.Context, in *pb.ValidateRequest, _ ...grpc.CallOption) (res *pb.ValidateResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.ValidateResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) RefreshToken(_ context.Context, in *pb.RefreshTokenRequest, _ ...grpc.CallOption) (res *pb.RefreshTokenResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*pb.RefreshTokenResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Validate(_ context.Context, token string) (payload *dto.TokenPayloadAuth, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		payload = args.Get(0).(*dto.TokenPayloadAuth)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return payload, err
}
