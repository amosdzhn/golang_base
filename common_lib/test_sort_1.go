package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func (p person) string() string {
	return fmt.Sprintf("%s:%d", p.name, p.age)
}

// ByAge 实现 sort.Interface 接口
type ByAge []person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func example() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
func main() {
	example()
}
