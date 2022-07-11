package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type sys_user struct {
	gorm.Model
	id              int
	uuid            string
	createTime      time.Time
	updateTime      time.Time
	createBy        string
	updateBy        string
	isDeleted       int
	accessKeyId     string
	accessKeySecret string
	role            int
}

func main() {
	dsn := "root:123456@tcp(127.0.0.2)/db_amosblog?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("db=%v \n", db)
}
