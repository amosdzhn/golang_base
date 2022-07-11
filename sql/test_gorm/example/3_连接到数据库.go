package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
连接到MySQL
GORM官方支持的数据库类型有:MySQL,PostgreSQL，SQlite，SQL，Server
*/
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_stu_gorm?charset=utf8mb4&parseTime=true&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//MySQL驱动程序提供了一些高级配置，可以在初始化过程中使用
	gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
}
