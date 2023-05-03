package dto

type ChatWebSocketType uint

const (
	ChatWebSocketUnknownType = iota
	ChatWebSocketLogin
	ChatWebSocketPrivateChat
	ChatWebSocketGroupChat
	ChatWebSocketChatMessage
	ChatWebSocketErrorMessage
)

type ChatWebSocketLoginDto struct {
	Type  ChatWebSocketType `json:"type" validate:"required"`
	Token string            `json:"token" validate:"jwt"`
}

type ChatWebSocketMessageDto struct {
	// Type can be both PrivateChat and GroupChat
	Type    ChatWebSocketType `json:"type" validate:"required"`
	Message string            `json:"message" validate:"required"`
	Targets []string          `json:"targets" validate:"uuid"`
	GroupID string            `json:"group_id" validate:"omitempty,uuid"`
}

type ChatWebSocketChatMessageDto struct {
	Type    ChatWebSocketType `json:"type"`
	Message string            `json:"message"`
}

type WebSocketErrorMessage struct {
	Type    ChatWebSocketType `json:"type"`
	Message string            `json:"message"`
	Data    any               `json:"data"`
}

type BroadcastEventDto struct {
	Message string   `json:"message"`
	UserId  string   `json:"user_id" validate:"uuid"`
	Targets []string `json:"targets" validate:"uuid"`
}
