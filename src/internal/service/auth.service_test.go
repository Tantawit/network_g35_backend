package service

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-chat/src/mock"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type AuthServiceTest struct {
	suite.Suite
	credential  *pb.Credential
	payload     *dto.TokenPayloadAuth
	authInfo    *dto.AuthInfo
	registerDto *dto.RegisterDto
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthServiceTest))
}

func (t *AuthServiceTest) SetupTest() {
	t.credential = &pb.Credential{
		AccessToken:  faker.Word(),
		RefreshToken: faker.Word(),
		ExpiresIn:    3600,
	}

	t.payload = &dto.TokenPayloadAuth{
		UserId: faker.UUIDDigit(),
	}

	t.authInfo = &dto.AuthInfo{
		Hostname:  faker.DomainName(),
		UserAgent: faker.Word(),
		IPAddress: faker.IPv4(),
	}

	t.registerDto = &dto.RegisterDto{
		Username:  faker.Username(),
		Password:  faker.Password(),
		Title:     faker.Word(),
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
	}
}

func (t *AuthServiceTest) TestValidateSuccess() {
	want := t.payload
	token := faker.Word()

	c := mocks.ClientMock{}
	c.On("Validate", &pb.ValidateRequest{Token: token}).Return(&pb.ValidateResponse{
		UserId: t.payload.UserId,
	}, nil)

	srv := AuthServiceImpl{
		Client: &c,
		Logger: gosdk.NewLogger("auth-test"),
	}

	actual, err := srv.Validate(context.Background(), token)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *AuthServiceTest) TestValidateGrpcErr() {
	token := faker.Word()

	c := mocks.ClientMock{}
	c.On("Validate", &pb.ValidateRequest{Token: token}).Return(nil, status.Error(codes.Unavailable, "service down'"))

	srv := AuthServiceImpl{
		Client: &c,
		Logger: gosdk.NewLogger("auth-test"),
	}

	actual, err := srv.Validate(context.Background(), token)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), status.Error(codes.Unavailable, "service down'"), err)
}
