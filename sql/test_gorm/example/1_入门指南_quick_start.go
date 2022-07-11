package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
gorm 快速入门
*/

// User 定义模型
type User struct {
	gorm.Model
	ID       int
	Username string
	Password string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.2)/db_stu_gorm?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}

	err = db.AutoMigrate(&User{ID: 1, Username: "张飞", Password: "aabbcc"})
	if err != nil {
		panic("fiald automigrate")
	}
	// create
	db.Create(&User{ID: 1, Username: "张三", Password: "pwdzs"})
	// read
	var user User
	db.First(&user, "id = ?", "1")
	fmt.Printf("user=%v \n", user)
	// update
	db.Model(&user).Update("username", "李四")
}
