package user

import "github.com/FuradWho/BlockchainDataColla/orgDeploy/internal/apiserver/store"

type UserController struct {
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{}
}
