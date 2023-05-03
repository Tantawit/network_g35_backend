package entity

import "github.com/google/uuid"

type ChatUserMembership struct {
	UserID  *uuid.UUID `json:"user_id"`
	GroupID *uuid.UUID `json:"group_id"`
}
