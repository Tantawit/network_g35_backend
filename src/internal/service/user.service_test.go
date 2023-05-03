package service

import (
	"context"
	mocks "github.com/2110336-2565-2/cu-freelance-chat/src/mock"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type UserServiceSuite struct {
	suite.Suite
	srv    UserServiceImpl
	client *mocks.UserClientMock

	user *pb.UserCUFreelance
}

func (t *UserServiceSuite) SetupTest() {
	logger := gosdk.NewLogger("user-service-test")
	t.client = &mocks.UserClientMock{}
	t.srv = UserServiceImpl{
		Logger: logger,
		Client: t.client,
	}

	t.user = &pb.UserCUFreelance{
		Id:          uuid.NewString(),
		Username:    gosdk.StringAdr(faker.Username()),
		Title:       faker.Word(),
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		Email:       faker.Email(),
		DisplayName: faker.Username(),
		Faculty:     gosdk.StringAdr(faker.Word()),
	}
}

func (t *UserServiceSuite) TestFindOneCUFreelanceSuccess() {
	t.client.On("FindOneUserCUFreelance", mock.AnythingOfType("*pb.FindOneUserCUFreelanceRequest")).
		Return(&pb.FindOneUserCUFreelanceResponse{User: t.user}, nil).
		Once()

	actual, err := t.srv.FindOneCUFreelance(context.Background(), t.user.Id)

	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), actual)
	assert.Equal(t.T(), t.user.Id, actual.Id)
	assert.Equal(t.T(), t.user.Firstname, actual.Firstname)
	assert.Equal(t.T(), t.user.Lastname, actual.Lastname)
	assert.Equal(t.T(), t.user.Email, actual.Email)
	assert.Equal(t.T(), t.user.Faculty, actual.Faculty)

	t.client.AssertExpectations(t.T())
}

func (t *UserServiceSuite) TestFindOneCUFreelanceNotFound() {
	t.client.On("FindOneUserCUFreelance", mock.AnythingOfType("*pb.FindOneUserCUFreelanceRequest")).
		Return(nil, status.Error(codes.NotFound, "not found")).
		Once()

	actual, err := t.srv.FindOneCUFreelance(context.Background(), t.user.Id)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())

	t.client.AssertExpectations(t.T())
}

func (t *UserServiceSuite) TestFindOneCUFreelanceInternalErr() {
	t.client.On("FindOneUserCUFreelance", mock.AnythingOfType("*pb.FindOneUserCUFreelanceRequest")).
		Return(nil, status.Error(codes.Internal, "internal error")).
		Once()

	actual, err := t.srv.FindOneCUFreelance(context.Background(), t.user.Id)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())

	t.client.AssertExpectations(t.T())
}

func (t *UserServiceSuite) TestFindOneCUFreelanceGrpcErr() {
	t.client.On("FindOneUserCUFreelance", mock.AnythingOfType("*pb.FindOneUserCUFreelanceRequest")).
		Return(nil, status.Error(codes.Unavailable, "service down")).
		Once()

	actual, err := t.srv.FindOneCUFreelance(context.Background(), t.user.Id)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())

	t.client.AssertExpectations(t.T())
}
