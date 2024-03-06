package common

import (
	"Panda/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库
var DB *gorm.DB

// NewDBConnection 新建数据库连接并返回初始化数据库
func NewDBConnection() *gorm.DB {
	driverName := util.GetString("datasource.driver_name")
	host := util.GetString("datasource.host")
	port := util.GetString("datasource.port")
	database := util.GetString("datasource.database")
	username := util.GetString("datasource.username")
	password := util.GetString("datasource.password")
	charset := util.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DriverName: driverName,
			DSN:        dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
		})
	if err != nil {
		panic("fail to connect databse,err:" + err.Error())
	}
	DB = db
	return db
}

// GetDB 获取数据库
func GetDB() *gorm.DB {
	return DB
}
