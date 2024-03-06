package entity

import "time"

// User 用户实体
type User struct {
	// UserId 主键id
	UserId int `gorm:"primary_key;column:user_id" json:"UserId"`
	// WechatAvatar 微信头像
	WechatAvatar string `gorm:"column:wechat_avatar" json:"WechatAvatar"`
	// WechatName 微信名称
	WechatName string `gorm:"column:wechat_name" json:"WechatName"`
	// OpenID 授权用户唯一标识
	OpenID string `gorm:"column:openid" json:"OpenID"`
	// UnionId 当且仅当该网站应用已获得该用户的 userinfo 授权时，才会出现该字段
	UnionId string `gorm:"column:unionid" json:"UnionId"`
	// Type 头像类型
	Type int `gorm:"column:type" json:"Type"`
	// IsAdmin 是否管理员
	IsAdmin int `gorm:"column:is_admin" json:"IsAdmin"`
	// IsEnabled 是否启用
	IsEnabled int `gorm:"column:is_enabled" json:"IsEnabled"`
	// CreateAtUtc  创建时间
	CreateAtUtc time.Time `gorm:"column:create_at_utc" json:"CreateAtUtc"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
