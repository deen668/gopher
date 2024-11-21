package main

import "fmt"

// const  // 全局常量
//func main() {
//	const (
//		a = (iota + 1) * (iota + 1) //iota=0,行索引，索引开发人员习惯一般从0开始
//		b
//		c
//		d
//	) // 局部常量，只能在当前函数main中使用
//	fmt.Println(a, b, c, d)
//}

func main() {
	const (
		a = 123
		b = (iota + 1) * (iota + 1)
		c
		d
	)
	fmt.Println(a, b, c, d)
}
