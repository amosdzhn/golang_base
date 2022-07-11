package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"op_mysql/db_mysql"
	"runtime"
	"time"
)

// 简单测试mysql Insert操作
func testInsert() {
	ret, err := db_mysql.DB.Exec("insert into sys_user values(?,?,?,?,?,?,?,?,?,?)",
		2, 222, time.Now(), time.Now(), "go_client", "go_client", 0, "b1", "b1", 1)
	if err != nil {
		fmt.Println(err.Error())
		runtime.Goexit()
	}
	id, err := ret.LastInsertId()
	affected, err := ret.RowsAffected()
	fmt.Printf("插入结果:%v 插入的id=%v 受影响的行数=%v \n", ret, id, affected)
}

// 简单测试mysql Query操作
type user struct {
	id                int
	uuid              string
	create_time       time.Time
	update_time       time.Time
	create_by         string
	update_by         string
	is_deleted        int
	access_key_id     string
	access_key_secret string
	role              int
}

func testQueryOne() {
	//单行查询
	queryRowSQL := "select * from sys_user where id = ?"
	row := db_mysql.DB.QueryRow(queryRowSQL, 1)
	aUser := new(user)
	err := row.Scan(&aUser.id, &aUser.uuid, &aUser.create_time, &aUser.update_time,
		&aUser.create_by, &aUser.update_by, &aUser.is_deleted, &aUser.access_key_id,
		&aUser.access_key_secret, &aUser.role)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("查询到的结果为: %v \n", *aUser)
	}
}

func main() {
	err := db_mysql.InitMySQL()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("连接成功")
	}

	// 测试插入
	//testInsert()

	// 测试查询单行
	testQueryOne()
}
