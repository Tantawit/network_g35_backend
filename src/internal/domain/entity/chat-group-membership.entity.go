package entity

import "github.com/google/uuid"

type ChatGroupMembership struct {
	GroupID *uuid.UUID `json:"group_id"`
	UserID  *uuid.UUID `json:"user_id"`
}
