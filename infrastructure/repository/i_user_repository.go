package repository

import (
	"Panda/infrastructure/entity"
)

// IUserRepository 用户仓储接口
type IUserRepository interface {
	Get(userId int) (*entity.User, error)
}
