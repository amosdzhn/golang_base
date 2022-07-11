package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
}

/**
GORM提供了First,Take,Last方法，以便从数据库从检索单个记录。
查询数据库时添加了LIMIT 1 的条件，且没有找到记录时会返回ErrRecordNotFound 错误
*/

//  db.First(&user) 获取第一条记录 SELECT * from users ORDER BY id LIMIT 1
// db.Take(&user) 获取一条记录，没有指定排序字段 SELECT * FROM users LIMIT 1
func testFirstTakeLast(db *gorm.DB) {
	var user Users
	tx := db.First(&user)
	fmt.Printf("first: tx.Error = %v \n", tx.Error)
	fmt.Printf("fitst: user = %v \n", user)

	db.Take(&user)
	fmt.Printf("take user = %v \n", user)

	db.Last(&user)
	fmt.Printf("last user = %v \n", user)
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db_stu_gorm?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("mysql 连接错误")
	}
	testFirstTakeLast(db)

}
