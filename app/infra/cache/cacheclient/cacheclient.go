package cacheclient

import (
	"context"
	"strconv"

	domainClient "github.com/Akito-Fujihara/web-application-template/app/domain/cacheclient"
	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var Set = wire.NewSet(
	NewCacheClient,
)

type CacheClient struct {
	Client *redis.Client
}

func NewCacheClient(client *redis.Client) domainClient.ICacheClient {
	return &CacheClient{
		Client: client,
	}
}

type sessionVlaue struct {
	ID string `redis:"id"`
	Name string `redis:"name"`
	Email string `redis:"email"`
}

func (c *CacheClient) SetSession(ctx context.Context, sessionID string, user *model.User) error {
	// 1日にセッションが削除されるように設定
	dto := buildUserToDTO(user)
	if err := c.Client.HSet(ctx, sessionID, dto).Err(); err != nil {
		return err
	}
	return c.Client.Expire(ctx, sessionID, 86400).Err()
}

func buildUserToDTO(user *model.User) *sessionVlaue {
	return &sessionVlaue{
		ID: strconv.FormatInt(user.ID, 10),
		Name: user.Name,
		Email: user.Email,
	}
}
