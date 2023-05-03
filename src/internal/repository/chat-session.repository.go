package repository

import gosdk "github.com/2110336-2565-2/cu-freelance-library"

type ChatSessionRepositoryImpl struct {
	Repo gosdk.RedisRepository
}

func (r *ChatSessionRepositoryImpl) FindServer(userId string) (string, error) {
	return r.Repo.GetHashCache(userId, "server-id")
}

func (r *ChatSessionRepositoryImpl) RegisterClient(userId string, serverId string) error {
	return r.Repo.SaveHashCache(userId, "server-id", serverId, -1)
}

func (r *ChatSessionRepositoryImpl) UnregisterClient(userId string) error {
	return r.Repo.RemoveCache(userId)
}
