package main

import "fmt"

//func main() {
//	//iota 定义成批常量，iota从0，从成批的第一行开始，每一行加1
//	const (
//		a = iota + 1
//		b //批量定义的时候，下面一行可以继承上面的表达式，相当于b=iota
//		c
//		_ //匿名变量，占位符
//		d
//		t = 100
//		_
//		_
//		e
//	)
//	fmt.Println(a, b, c, d, e)
//}

//func main() {
//	const (
//		m = 100
//		n
//		a = iota
//		b
//		c
//		d
//	)
//	fmt.Println(m, n, a, b, c, d)
//	//const a = [2]int{1, 2} //数组，长度是常量，数组不能做常量,[2]int是类型，底层实现有关，数组值是不可变的
//	var a1 = [2]int{1, 2}
//	fmt.Println(a1)
//}

const ( //全局作用域，不在函数内部，全局可见
	m = 100
	n
	a = iota
)

func main() {
	fmt.Println(m, n, a)
}
