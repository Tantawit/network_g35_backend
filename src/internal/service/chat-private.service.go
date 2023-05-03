package service

import (
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/entity"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ChatPrivateServiceImpl struct {
	Logger                gosdk.Logger
	ChatPrivateRepository repository.ChatPrivateRepository
}

func (s *ChatPrivateServiceImpl) SendMessage(chat *entity.ChatPrivate) (bool, error) {
	if err := s.ChatPrivateRepository.Create(chat); err != nil {
		s.Logger.Error(err).
			Keyword("user_id", chat.User1ID).
			Msg(err.Error())
		return false, status.Error(codes.Internal, err.Error())
	}

	s.Logger.Info().
		Keyword("message_id", chat.MessageID).
		Keyword("user1_id", chat.User1ID).
		Keyword("user2_id", chat.User2ID).
		Msg("chat message saved successfully")

	return true, nil
}

func (s *ChatPrivateServiceImpl) ReadMessage(user1Id string, user2Id string, messageId string) (bool, error) {
	if err := s.ChatPrivateRepository.Read(user1Id, user2Id, messageId); err != nil {
		s.Logger.Error(err).
			Keyword("user_id", user1Id).
			Msg(err.Error())
		return false, status.Error(codes.Internal, err.Error())
	}

	s.Logger.Info().
		Keyword("message_id", messageId).
		Keyword("user1_id", user1Id).
		Keyword("user2_id", user2Id).
		Msg("chat message saved successfully")

	return true, nil
}

func (s *ChatPrivateServiceImpl) FindAllRecentChat(userId string) ([]*entity.ChatPrivate, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatPrivateServiceImpl) FindMessageSince(metadata *gosdk.PaginationMetadata, user1ID string, user2ID string, sentAt *time.Time) ([]*entity.ChatPrivate, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatPrivateServiceImpl) FindMessage(metadata *gosdk.PaginationMetadata, user1ID string, user2ID string) ([]*entity.ChatPrivate, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ChatPrivateServiceImpl) Sync(metadata *gosdk.PaginationMetadata, user1ID string, user2ID string, syncToken string) ([]*entity.ChatPrivate, error) {
	//TODO implement me
	panic("implement me")
}
