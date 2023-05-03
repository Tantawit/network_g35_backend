package entity

import (
	"github.com/google/uuid"
	"time"
)

type ChatGroup struct {
	GroupID   *uuid.UUID `json:"group_id"`
	UserID    *uuid.UUID `json:"user_id"`
	MessageID *uuid.UUID `json:"message_id"`
	SentAt    time.Time  `json:"sent_at"`
	ReadAt    time.Time  `json:"read_at"`
	Message   string     `json:"message"`
}
