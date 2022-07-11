package main

import (
	"fmt"
	"sort"
)

/**
sort包提供了对切片和用户自定义数据集的排序等函数
type Interface interface{
	Len() int //返回集合中元素个数
	Less(i,j int) bool //索引i元素的值是否比索引j元素的值小
	Swap(i,j int) //交换i,j位置的元素
}
type IntSlice
type Float64Slice
type StringSlice
*/
type personSlice []person

type person struct {
	name string
	age  int
}

func (ps personSlice) Len() int {
	return len(ps)
}

func (ps personSlice) Less(i, j int) bool {
	return ps[i].age > ps[j].age
}

func (ps personSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

//func (ps personSlice) Search(age int) person {
//
//}

func main() {
	// 1.创建一个整型切片
	var s1 = []int{11, 5, 25, 19, 9}
	// 2.给整型切片进行递增排序
	sort.Ints(s1)
	// 3.打印结果
	fmt.Printf("结果:%v \n", s1)
	// 4.查看一个切片是否是递增的
	s1IsSorted := sort.IntsAreSorted(s1)
	fmt.Printf("是否是递增的:%v\n", s1IsSorted)

	// 1.定义一个IntSlice类型的变量
	s2 := sort.IntSlice(s1)
	fmt.Printf("%T %T \n", s1, s2)

	// 1.创建一个personSlice类型的切片
	pss := personSlice{
		{name: "张飞", age: 18},
		{name: "刘备", age: 20},
		{name: "关羽", age: 19},
	}
	// 2.调用方法
	fmt.Printf("切片的元素个数:%v 张飞比刘备年纪大吗?:%v\n", pss.Len(), pss.Less(0, 1))
	pss.Swap(0, 1)
	fmt.Printf("交换张飞和刘备在切片中的位置后的切片:%v\n", pss)
	// 3.对整个切片进行排序
	sort.Sort(pss)
	fmt.Printf("切片排序之后的结果:%v\n", pss)

}
