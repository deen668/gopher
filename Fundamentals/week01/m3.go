package main

import "fmt"

func main() {
	// iota定义成批变量
	//const (
	//	a = iota + 1
	//	b // 批量定义的时候，下面一行可以继承上面的表达式，b = iota
	//	c
	//	_ // 空白标识符，用于占位
	//	d
	//	t = 100
	//	_
	//	_
	//	e
	//)
	//fmt.Println(a, b, c, d, t, e)

	// iota定义成批常量,iota从0开始，从成批的第一行开始，每次递增1
	const (
		m = 100
		n
		a = iota * 2
		b
		c
		d
	)
	fmt.Println(m, n, a, b, c, d)
	//const m1 = [2]int{100,200} //[2]int类型，代表的是数组,底层实现有关
	var m1 = [2]int{100, 200}
	fmt.Println(m1)
}
