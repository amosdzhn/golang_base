package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
测试 create 记录

1.直接创建单条记录
2.使用指定的字段创建记录
3.忽略指定的字段创建记录
4.批量一次性创建记录
5.批量多次创建记录
*/

type Users struct {
	gorm.Model
	Username string
	Password string
}

// 创建记录
func create(db *gorm.DB) {
	user := Users{Username: "张海客", Password: "123456"}
	//创建记录(通过数据的指针来创建)
	result := db.Create(&user)
	fmt.Printf("返回插入数据的主键=%v 返回的错误=%v 返回插入记录的条数=%v", user.ID, result.Error, result.RowsAffected)
}

// 用指定的字段创建记录
func selectCreate(db *gorm.DB) {
	user := Users{Username: "王非", Password: "123456"}
	db.Select("Username").Create(&user)
}

// 忽略指定的字段进行创建记录
func omitCreate(db *gorm.DB) {
	user := Users{Username: "诸葛亮", Password: "123aaa"}
	db.Omit("Username").Create(&user)
}

// 批量插入
// 要有效的插入大量记录，需要将一个slice 传递给create方法。
// GORM会生成一条单独的SQL语句来插入所有的数据，并回填所有主键的值，钩子方法也会被调用
func bathCreate(db *gorm.DB) {
	users := []Users{
		{Username: "bath1", Password: "bathpwd1"},
		{Username: "bath2", Password: "bathpwd2"},
		{Username: "bath3", Password: "bathpwd3"},
	}
	db.Create(&users)

	for _, user := range users {
		fmt.Println("新插入的主键=", user.ID)
	}
}

func createInBatches(db *gorm.DB) {
	users := []Users{
		{Username: "bath1", Password: "bathpwd1"},
		{Username: "bath2", Password: "bathpwd2"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
		{Username: "bath3", Password: "bathpwd3"},
	}
	db.CreateInBatches(&users, 3)

	for _, user := range users {
		fmt.Printf("生成的主键:%v \n", user.ID)
	}
}

// 创建钩子
// GORM允许用户定义的钩子有: BeforeSave,BeforeCreate,AfterSave,AfterCreate 创建记录时将调用这些钩子方法
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("execute BeforeCreate Hook")
	return nil
}

// 根据Map创建记录
// GORM支持 map[string]interface{} 和 []map[string]interface{} 创建记录
// 根据map创建记录时，association不会被调用,并且主键也不会自动填充
func mapCreate(db *gorm.DB) {
	// 单个create
	mapUser := map[string]interface{}{
		"Username": "map_name_1",
		"Password": "map_pwd_1",
	}
	db.Model(&Users{}).Create(mapUser)
	for k, v := range mapUser {
		fmt.Printf("k=%v v=%v \n", k, v)
	}
}
func sliceMapCreate(db *gorm.DB) {
	sliceMapUsers := []map[string]interface{}{
		{"Username": "slicemap_uname_1", "Password": "slicemap_pwd_1"},
		{"Username": "slicemap_uname_2", "Password": "slicemap_pwd_2"},
		{"Username": "slicemap_uname_3", "Password": "slicemap_pwd_3"},
	}
	tx := db.Model(&Users{}).Create(sliceMapUsers)
	fmt.Printf("err=%v rowaf=%v \n", tx.Error, tx.RowsAffected)
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/db_stu_gorm?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("mysql 连接错误")
	}
	//create(db)

	//selectCreate(db)

	//omitCreate(db)

	//bathCreate(db)

	//createInBatches(db)

	//mapCreate(db)

	sliceMapCreate(db)
}
