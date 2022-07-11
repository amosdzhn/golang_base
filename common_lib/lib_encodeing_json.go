package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// encoding/json 包 包含了json转换为struct或struct转换为json的功能
/**
两个核心函数：
Marshal(v interface{}) ([]byte,error) 将一个任意类型的变量转换为字节数组(字节切片可转换为字符串)
Unmarshal(data []byte,v interface{}) error 将字节切片转换为变量
两个核心结构体：
type Decoder struct{...} 从输入流读取并解析为JSON
type Encoder struct{...} 写JSON到输出流
*/

// 提示：结构体属性必须是public
type student struct {
	Sid  int
	Name string
	Age  int
}

func main() {
	// 测试struct -> JSON
	student1 := student{Sid: 1, Name: "张海客", Age: 20}
	stuJSON, err := json.Marshal(student1)
	if err != nil {
		runtime.Goexit()
	}
	fmt.Printf("student_JSON= %v\n", string(stuJSON))
	v1, err := json.Marshal(1)
	if err != nil {
		fmt.Println(err.Error())
		runtime.Goexit()
	}
	fmt.Printf("v1= %v\n", string(v1))

	student2 := map[string]interface{}{
		"name": "张飞",
		"age":  10,
	}
	student2JSON, err := json.Marshal(student2)
	if err != nil {
		fmt.Println(err.Error())
		runtime.Goexit()
	}
	fmt.Printf("student2_JSON= %v\n", string(student2JSON))

	//测试JSON -> struct
	aJsonStr := "{\"sid\":1001,\"name\":\"张飞\",\"age\":20}"
	var student3 = new(student)
	fmt.Printf("%v %v %v %v\n", student3.Sid, student3.Name, student3.Age, &student3)
	err = json.Unmarshal([]byte(aJsonStr), student3)
	if err != nil {
		fmt.Println(err.Error())
		runtime.Goexit()
	}
	fmt.Printf("%v %v %v %v\n", student3.Sid, student3.Name, student3.Age, &student3)
}
