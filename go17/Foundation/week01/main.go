package main

import (
	"fmt"
)

// func main() {
// 	a := 100
// 	fmt.Println("Hello, World!")
// 	fmt.Println("magedu.com", a)
// }
/* 多行注释
x int
y int
returns: int
函数说明
*/
// func add(x, y int) int {
// 	// 这里写下一句或下几句的注释
// 	return x + y // 本行注释，如果短，可以写这里，如果长，写上面
// }

// 函数注释也可以这样多行
// 写上面
// func main() {
// 	fmt.Println(add(4, 5))
// 	// TODO: 之后完成某某功能
// 	// NOTE: 请注意，该程序只可使用开源
// 	// Deprecated: 告知已经过期，建议不要使用。未来某个版本可以移除
// 	fmt.Println(100, 1_0000) // 1_0000=1万
// 	fmt.Println(0x6162, 0x61_62_63)
// 	fmt.Println(0b0101, 0o101)
// 	fmt.Println(3.14)
// 	fmt.Println(3.14e2)
// 	fmt.Println(3.14e-2)
// 	fmt.Println('测')
// 	fmt.Printf("%T,%d,%c", '测', '测', '测')
// 	fmt.Println('\u6d4b')
// 	fmt.Println('\x31')
// 	fmt.Println('1')
// 	fmt.Println('\n')
// 	fmt.Println("abc", "\x61b\x63")
// 	fmt.Println("测试", "\u6d4b试")
// 	fmt.Println("n")
// 	fmt.Println(true, false)
// }

// 赋值等式右边用“无类型常量untyped constant”来赋值
// const a int = 100 // 指定类型定义常量并赋值
//	func main() {
//		const ( // 定义常量，等式左边未给出类型，将进行类型推导
//			b = "abc"
//			c = 12.3
//			d = 'T'
//		)
//		fmt.Println(a, b, c, d)
//	}
//
// iota
//
//	func main() {
//		// 单独写iota从0开始，iota不用在这里
//		const a = iota
//		const b = iota
//		fmt.Println(a, b)
//	}
//
// iota常用在常量批量定义中，是写在批量定义的括号里的，以定义星期常量为例
// func main() {
// 	const (
// 		SUN = iota // 0
// 		MON = iota // 1
// 		TUE = iota // 2
// 	)
// 	fmt.Println(SUN, MON, TUE)
// }

//	func main() {
//		const (
//			SUN = iota // 0
//			MON
//			TUE
//		)
//		fmt.Println(SUN, MON, TUE)
//	}
//
// 比较繁琐的写法，仅作测试
//
//	func main() {
//		const (
//			a = iota
//			b
//			c
//			_
//			d
//			e = 10
//			f
//			g = iota
//			h
//		)
//		fmt.Println(a, b, c, d, e, f, g, h)
//	}
//
// 批量写iota从0开始，智能重复上一行公式
// func main() {
// 	const (
// 		a = 2 * iota
// 		b
// 		c
// 		d
// 	)
// 	fmt.Println(a, b, c, d)
// }

// func main() {
// 	const (
// 		m = 0
// 		n
// 		a = 2 * iota
// 		b
// 		c
// 		d
// 	)
// 	fmt.Println(a, b, c, d, m, n)
// }

// func main() {
// 	var name string
// 	var age int
// 	var isOK bool
// 	fmt.Println(name, age, isOK)

// 	var (
// 		a string
// 		b int
// 		c bool
// 		d float32
// 	)
// 	fmt.Println(a, b, c, d)
// }

// func main() {
// 	// var name string = "七米"
// 	// var age int = 18
// 	name, age := "七米", 18
// 	fmt.Println(name, age)
// }

// func main() {
// 	age := 20
// 	fmt.Println(age)
// }

// func foo() (int, string) {
// 	return 18, "七米"
// }

// func main() {
// 	x, _ := foo()
// 	_, y := foo()
// 	fmt.Println("x=", x)
// 	fmt.Println("y=", y)
// }

// func main() {
// 	// const pi = 3.1415
// 	// const e = 2.7182
// 	const (
// 		pi = 3.1415
// 		e  = 2.7182
// 	)
// 	fmt.Println(pi, e)
// }

// func main() {
// 	const (
// 		n1 = 100
// 		n2
// 		n3
// 	)
// 	fmt.Println(n1, n2, n3)
// }

// func main() {
// 	const (
// 		leve10 = iota
// 		leve11
// 		leve12
// 		leve13
// 	)
// 	fmt.Println(leve10, leve11, leve12, leve13)
// }

// func main() {
// 	const (
// 		_  = iota
// 		KB = 1 << (10 * iota) // 1<<（10*1)
// 		MB = 1 << (10 * iota) // 1<<（10*2)
// 		GB = 1 << (10 * iota) // 1<<（10*3)
// 		TB = 1 << (10 * iota) // 1<<（10*4)
// 		PT = 1 << (10 * iota) // 1<<（10*5)
// 	)
// 	fmt.Println(KB, MB, GB, TB, PT)
// }

func main() {
	const (
		a, b = iota + 1, iota + 2
		c, d
		e, f
	)
	fmt.Println(a, b, c, d, e, f)
}
