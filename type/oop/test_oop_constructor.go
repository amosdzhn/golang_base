package main

import "fmt"

/**
go语言中没有构造方法的概念，但是我们可以使用普通方法模拟一个构造方法
*/

type book struct {
	name  string
	price float32
}

func (b *book) printBook() {
	fmt.Println(*b)
}

func newBook(name string, price float32) (*book, error) {
	if name == "" {
		return nil, fmt.Errorf("书名不能为:%v", name)
	}
	if price < 0 {
		return nil, fmt.Errorf("价格不能为负数:%v", price)
	}
	return &book{
		name,
		price,
	}, nil
}
func main() {
	b, err := newBook("数据结构", 20.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*b)
}
