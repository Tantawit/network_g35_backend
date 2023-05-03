package entity

import (
	"github.com/google/uuid"
	"time"
)

type ChatPrivate struct {
	User1ID   *uuid.UUID `json:"user1_id"`
	User2ID   *uuid.UUID `json:"user2_id"`
	MessageID *uuid.UUID `json:"message_id"`
	SentAt    time.Time  `json:"sent_at"`
	ReadAt    time.Time  `json:"read_at"`
	Message   string     `json:"message"`
}
