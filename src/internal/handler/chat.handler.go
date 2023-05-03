package handler

import (
	"context"
	"encoding/json"
	"github.com/2110336-2565-2/cu-freelance-chat/src/internal/domain/dto"
	"github.com/2110336-2565-2/cu-freelance-chat/src/metric"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/validator"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/websocket/v2"
	"time"
)

type IChatContext interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(messageType int, data []byte) error
	GetConn() *websocket.Conn
	Close() error
}

type ChatHandlerImpl struct {
	Logger     gosdk.Logger
	Unregister chan *websocket.Conn
	Broadcast  chan []byte
	AuthSvc    service.AuthService
	UserSvc    service.UserService
	//ChatPrivateSvc    service.ChatPrivateService
	//ChatGroupSvc      service.ChatGroupService
	WebsocketSvc      service.WebsocketService
	KeepAliveInterval int
	Validator         validator.Validator
}

func (h *ChatHandlerImpl) Listen(ctx IChatContext) {
	var (
		msg []byte
		err error
	)

	defer func() {
		h.Unregister <- ctx.GetConn()
		ctx.Close()
	}()

	if _, msg, err = ctx.ReadMessage(); err != nil {
		if !websocket.IsCloseError(err, websocket.CloseNoStatusReceived, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
			sentry.CaptureException(err)
			h.Logger.Error(err).
				Msg(err.Error())
		}
		return
	}

	// Validate user identity
	loginDto := &dto.ChatWebSocketLoginDto{}
	if err := json.Unmarshal(msg, loginDto); err != nil {
		h.sendErrMsg(ctx, "invalid json format", nil)
		h.Logger.
			Warn().
			Keyword("error", err).
			Msg("invalid json format")
		return
	}

	if err := h.Validator.Validate(loginDto); err != nil {
		h.sendErrMsg(ctx, "invalid token format", err)
		h.Logger.
			Warn().
			Keyword("error", err).
			Msg("invalid token format")
		return
	}

	if loginDto.Type != dto.ChatWebSocketLogin {
		h.sendErrMsg(ctx, "invalid message type", err)
		h.Logger.
			Warn().
			Keyword("error", err).
			Msg("invalid message type")
		return
	}

	tokenPayload, err := h.AuthSvc.Validate(context.Background(), loginDto.Token)
	if err != nil {
		h.Logger.Error(err).
			Msg(err.Error())
		return
	}

	user, err := h.UserSvc.FindOneCUFreelance(context.Background(), tokenPayload.UserId)
	if err != nil {
		h.Logger.Error(err).
			Msg(err.Error())
		return
	}

	res, _ := json.Marshal(map[string]bool{
		"connect_success": true,
	})

	if err := ctx.WriteMessage(websocket.TextMessage, res); err != nil {
		sentry.CaptureException(err)
		h.Logger.Error(err).
			Msg(err.Error())
		return
	}

	// Register user to mapping service
	if err := h.WebsocketSvc.RegisterClient(ctx.GetConn(), user.Id); err != nil {
		h.Logger.
			Warn().
			Keyword("user_id", user.Id).
			Keyword("display_name", user.DisplayName).
			Keyword("error", err.Error()).
			Msg("failed to register user to service")
	}

	h.Logger.Info().
		Keyword("user_id", user.Id).
		Keyword("display_name", user.DisplayName).
		Msg("user connected to server")

	metric.OnlineUserCounter.WithLabelValues().Inc()

	go h.keepAlive(ctx, user)

	for {
		messageType, message, err := ctx.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				h.Logger.
					Error(err).
					Keyword("user_id", user.Id).
					Msg("read error")
			}
			return
		}

		if messageType == websocket.TextMessage {
			h.Broadcast <- message
		}
	}
}

func (h *ChatHandlerImpl) keepAlive(ctx IChatContext, user *pb.UserCUFreelance) {
	for {
		timeout := make(chan struct{})
		go func() {
			time.Sleep(time.Second * time.Duration(h.KeepAliveInterval))
			timeout <- struct{}{}
		}()

		select {
		case <-timeout:
			if err := ctx.WriteMessage(websocket.PingMessage, []byte("keepalive")); err != nil {
				h.Logger.
					Warn().
					Keyword("user_id", user.Id).
					Keyword("display_name", user.DisplayName).
					Keyword("error", err.Error()).
					Msg("failed to send keep alive packet")
			}

		case msg := <-h.Broadcast:
			chatDto := &dto.ChatWebSocketMessageDto{}
			if err := json.Unmarshal(msg, chatDto); err != nil {
				h.Logger.
					Warn().
					Keyword("user_id", user.Id).
					Keyword("display_name", user.DisplayName).
					Keyword("error", err.Error()).
					Msg("failed to unmarshal chat message")
				continue
			}

			// Broadcast in websocket
			if err := h.WebsocketSvc.Broadcast(chatDto.Message, user.Id, chatDto.Targets); err != nil {
				h.Logger.
					Warn().
					Keyword("user_id", user.Id).
					Keyword("display_name", user.DisplayName).
					Keyword("error", err.Error()).
					Msg("failed to broadcast message")
			}

			// TODO using the chat service match to message type
			// Store message in persistence at Cassandra
			switch chatDto.Type {
			case dto.ChatWebSocketPrivateChat:
				// TODO call private chat service
			case dto.ChatWebSocketGroupChat:
				// TODO call group chat service
			default:
				h.Logger.
					Warn().
					Keyword("user_id", user.Id).
					Keyword("display_name", user.DisplayName).
					Msg("invalid chat type")
			}

		case _ = <-h.Unregister:
			metric.OnlineUserCounter.WithLabelValues(user.GetYear(), user.GetFaculty()).Dec()
			h.Logger.
				Info().
				Keyword("user_id", user.Id).
				Keyword("display_name", user.DisplayName).
				Msg("user left server")

			if err := h.WebsocketSvc.UnregisterClient(user.Id); err != nil {
				h.Logger.
					Warn().
					Keyword("user_id", user.Id).
					Keyword("display_name", user.DisplayName).
					Keyword("error", err.Error()).
					Msg("failed to unregister user from service")
			}
			return
		}
	}
}

func (h *ChatHandlerImpl) sendErrMsg(ctx IChatContext, msg string, data any) {
	errMsg := &dto.WebSocketErrorMessage{
		Type:    dto.ChatWebSocketErrorMessage,
		Message: msg,
		Data:    data,
	}

	raw, _ := json.Marshal(errMsg)

	if err := ctx.WriteMessage(websocket.TextMessage, raw); err != nil {
		h.Logger.
			Warn().
			Keyword("error", err.Error()).
			Msg("failed to send error message")
	}
}
