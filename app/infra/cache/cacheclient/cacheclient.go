package cacheclient

import (
	"context"
	"fmt"
	"time"

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
	ID int64 `redis:"id"`
	Name string `redis:"name"`
	Email string `redis:"email"`
}

func (c *CacheClient) GetSession(ctx context.Context, sessionID string) (*model.User, error) {
	var sv sessionVlaue
	if err := c.Client.HGetAll(ctx, sessionID).Scan(&sv); err != nil {
		return nil, err
	}

	// Session Expired
	if sv.ID == 0 {
		return nil, fmt.Errorf("session expired")
	}

	return buildDTOtoModel(&sv), nil
}

func (c *CacheClient) SetSession(ctx context.Context, sessionID string, user *model.User) error {	
	if err := c.Client.HSet(ctx, sessionID, buildModelToMap(user)).Err(); err != nil {
		return err
	}
	return c.Client.Expire(ctx, sessionID, 86400*time.Second).Err()
}

func buildModelToMap(user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
	}
}

func buildDTOtoModel(sv *sessionVlaue) *model.User {
	return &model.User{
		ID: sv.ID,
		Name: sv.Name,
		Email: sv.Email,
	}
}
