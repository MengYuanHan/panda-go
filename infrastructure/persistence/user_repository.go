package persistence

import (
	"Panda/common"
	"Panda/infrastructure/entity"
	"Panda/infrastructure/repository"
	_ "github.com/go-sql-driver/mysql" // driver
	"gorm.io/gorm"
)

// UsersRepository  实现仓储数据库连接
type UsersRepository struct {
	db *gorm.DB
}

// NewUserRepository 实现初始化仓库
func NewUserRepository() repository.IUserRepository {
	return &UsersRepository{db: common.GetDB()}
}

// Get 获取用户仓储
func (r *UsersRepository) Get(userId int) (*entity.User, error) {
	var users *entity.User
	usersSql := r.db.Table((*entity.User)(nil).TableName()).Where("user_id", userId)
	var err = usersSql.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
