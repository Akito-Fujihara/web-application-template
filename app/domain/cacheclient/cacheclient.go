package cacheclient

import (
	"context"

	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
)

type ICacheClient interface {
	SetSession(ctx context.Context, sessionID string, user *model.User) error
}
