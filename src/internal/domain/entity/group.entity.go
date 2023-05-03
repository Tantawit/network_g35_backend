package entity

import "github.com/google/uuid"

type Group struct {
	GroupID     *uuid.UUID `json:"group_id"`
	Name        string     `json:"name"`
	TotalMember int        `json:"total_member"`
}
