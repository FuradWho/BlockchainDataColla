package v1

import (
	"context"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/internal/apiserver/store"
	v1 "github.com/FuradWho/BlockchainDataColla/orgDeploy/internal/pkg/model/apiserver/v1"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *v1.User) error
	Update(ctx context.Context, user *v1.User) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) (*v1.User, error)
	List(ctx context.Context) (*v1.UserList, error)
}
type userService struct {
	store store.Factory
}

func (u userService) Create(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Update(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Delete(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Get(ctx context.Context, username string) (*v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) List(ctx context.Context) (*v1.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}
