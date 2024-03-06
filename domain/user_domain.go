package domain

import (
	"Panda/infrastructure/entity"
	"Panda/infrastructure/persistence"
	"Panda/infrastructure/repository"
	"context"
)

// UserDomain 用户应用领域
type UserDomain struct {
	Repository repository.IUserRepository
	Context    context.Context
}

// NewUserDomain 用户服务实例化
func NewUserDomain(ctx context.Context) IUserDomain {
	var userDomain UserDomain
	userDomain.Repository = persistence.NewUserRepository()
	userDomain.Context = ctx
	return &userDomain
}

// Get 根据id获取用户
func (userDomain *UserDomain) Get(userId int) (*entity.User, error) {
	return userDomain.Repository.Get(userId)
}
