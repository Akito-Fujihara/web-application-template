package private

import (
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/privateapi"
	"github.com/Akito-Fujihara/web-application-template/app/usecase"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewPrivateServer,
)

type PrivateServer struct{
	todoUsercase usecase.ITodoUsecase	
}

func NewPrivateServer(todoUsercase usecase.ITodoUsecase) privateapi.ServerInterface {
	return &PrivateServer{
		todoUsercase: todoUsercase,
	}
}

var _ privateapi.ServerInterface = (*PrivateServer)(nil)
