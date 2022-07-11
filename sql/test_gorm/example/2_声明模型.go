package main

import (
	"gorm.io/gorm"
	"time"
)

/**
1.模型定义
gorm中模型是标准的struct,由Go的基本数据类型，实现了Scanner和Valuer接口的自定义类型及其指针或别名组成。

2.约定
GORM倾向于约定而不是配置。默认情况下，GORM使用ID作为主键，使用结构体名的驼峰复数作为表名，字段名的
驼峰作为列名，并使用CreateAt,UpdateAt字段追踪创建和更新时间。

3.gorm.Model
GORM定义了一个gorm.Model结构体，其包括ID,CreateAt,UpdateAt,DeleteAt字段。我们可以将其镶嵌进结构体中，以包含这几个字段。

4.高级选项
可导出的字段在使用GORM进行CRUD时拥有全部的权限，此外，GORM还允许使用标签控制字段级别的权限。
这样我们可以让一个字段是只读，只写，只创建，只更新或被忽略。
注意：GORM Migrator进行创建表时，不会创建被忽略的字段
4.1字段级访问权限
4.2创建/更新时间追踪
4.3嵌入结构体
4.4字段标签
4.5关联标签

字段标签是可选的，GORM支持很多标签，标签大小写不敏感，建议使用驼峰
column 指定db列名
type 列数据类型
serializer 指定序列化程序如何序列化和反序列化到数据库
size 指定db列的长度大小
primaryKey 指定列作为主键
unique 指定列作为unique
default 指定列的默认值
precision 指定列的精度
scale 指定列的scale
not null 指定列为非空
autoIncrement 指定列自增
autoIncrementIncrement
embedded 嵌套字段
embeddedPrefix 嵌套字段前缀
autoCreateTime 创建时间追踪
autoUpdateTime 更新时间追踪
index 创建索引，多个名字相同创建复合索引
uniqueIndex 唯一索引,创建时唯一
check 创建约束
<- 设置字段写权限，<-:create表示只能创建 <-:update只能更新 <-:false没有写权限 <- 表示既可以create也可以update
-> 设置字段读权限,->:false 表示没有读权限
- 忽略这个字段，- 表示读写权限的时候忽略, -:migration 表示migration时忽略 -:all 表示没有读写和migration忽略
comment 当migration时添加注释
*/

type Teacher struct {
	Name string
}

type Student struct {
	//字段级访问权限
	Name  string `gorm:"<-:create"`          //allow read and create
	Name2 string `gorm:"<-:update"`          //allow read and update
	Name3 string `gorm:"<-"`                 // allow read and write(create and update)
	Name4 string `gorm:"<-:false"`           // allow read,disable write permission
	Name5 string `gorm:"->;<-:create"`       // allow read and create
	Name6 string `gorm:"->:false;<-:create"` //create only disable read from db
	Name7 string `gorm:"-"`                  // ignore this field when write and read with struct
	Name8 string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
	Name9 string `gorm:"-:migration"`        // ignore this field when migrate with struct

	//创建/更新时间追踪
	//GORM框架默认使用CreateAt UpdateAt 追踪创建和更新时间。如果定义了这两个字段，GORM在创建和更新时会自动填充当前时间
	//要使用不同名称的字段，您可以配置 autoCreateTime、autoUpdateTime 标签
	//如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int 即可
	CreateAt time.Time
	UpdateAt int
	Updated  int64 `gorm:"autoUpdateTime:nano"`  //在更新时使用纳秒时间戳填充
	Updated2 int64 `gorm:"autoUpdateTime:milli"` //在更新时使用毫秒时间戳填充
	Created  int64 `gorm:"autoCreateTime"`       //在创建时使用秒时间戳填充

	//嵌入结构体
	//对于匿名字段，GORM会将其字段包含在父结构当中
	gorm.Model
	//正常字段可使用 embedded 将其嵌入
	Teacher Teacher `gorm:"embedded"`
	//还可以使用 embeddedPrefix 来为db中的字段添加前缀,等效于将该类型的公有字段添加了前缀
	Teacher2 Teacher `gorm:"embedded;embeddedPredix:t_"`
}

func main() {

}
