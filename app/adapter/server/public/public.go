package public

import (
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
	"github.com/Akito-Fujihara/web-application-template/app/usecase"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewPublicServer,
)

type PublicServer struct{
	accountUsecase usecase.IAccountUsecase
}

func NewPublicServer(accountUsecase usecase.IAccountUsecase) publicapi.ServerInterface {
	return &PublicServer{
		accountUsecase: accountUsecase,
	}
}

var _ publicapi.ServerInterface = (*PublicServer)(nil)
