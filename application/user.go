package application

import (
	"Panda/domain"
	"Panda/infrastructure/entity"
	"context"
)

// UserInteractor 用户应用交互逻辑
type UserInteractor struct {
	userDomain domain.IUserDomain
	Context    context.Context
}

// NewUserInteractor 用户服务实例化
func NewUserInteractor(ctx context.Context) *UserInteractor {
	var users UserInteractor
	users.userDomain = domain.NewUserDomain(ctx)
	users.Context = ctx
	return &users
}

// Get 获取用户
func (srv *UserInteractor) Get(userId int) (*entity.User, error) {
	return srv.userDomain.Get(userId)
}
