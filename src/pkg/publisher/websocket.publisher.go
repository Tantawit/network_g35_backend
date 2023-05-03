package publisher

type WebsocketPublisher interface {
	RegisterClient(userId string, serverId string) error
	UnregisterClient(userId string) error
}
