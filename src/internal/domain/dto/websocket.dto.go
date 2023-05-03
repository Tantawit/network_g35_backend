package dto

type RegisterClientEventDto struct {
	UserID   string `json:"user_id"`
	ServerID string `json:"server_id"`
}

type UnregisterClientEventDto struct {
	UserID string `json:"user_id"`
}
