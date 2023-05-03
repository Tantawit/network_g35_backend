package dto

import "github.com/2110336-2565-2/cu-freelance-chat/src/constant"

type TokenPayloadAuth struct {
	UserId   string            `json:"user_id"`
	UserType constant.UserType `json:"user_type"`
}

type VerifyTicket struct {
	Ticket     string              `json:"ticket" validate:"required"`
	University constant.University `json:"university" validate:"required"`
}

type RegisterDto struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Title       string `json:"title"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"email"`
	Phone       string `json:"phone" validate:"phone_number"`
	DisplayName string `json:"display_name" validate:"required"`
}

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"password"`
}

type Validate struct {
	Token string `json:"token" validate:"jwt"`
}

type RedeemNewToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AuthInfo struct {
	Hostname  string `json:"hostname"`
	UserAgent string `json:"user_agent"`
	IPAddress string `json:"ip_address"`
}

type ChangePasswordDto struct {
	Username        string `json:"username"  validate:"required"`
	CurrentPassword string `json:"current_password"  validate:"password"`
	NewPassword     string `json:"new_password"  validate:"password"`
}
