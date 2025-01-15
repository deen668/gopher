package main

import "fmt"

func main() {
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
