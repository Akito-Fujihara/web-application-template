package cacheclient

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
)

type ICacheClient interface {
	GetSession(ctx context.Context, sessionID string) (*model.User, error)
	SetSession(ctx context.Context, sessionID string, user *model.User) error
}
