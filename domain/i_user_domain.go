package domain

import "Panda/infrastructure/entity"

// IUserDomain 用户领域接口
type IUserDomain interface {
	Get(userId int) (*entity.User, error)
}
