package main

import "fmt"

// var g int = 100 // 只能做到包内全局使用

// const e = 200

// type Student struct{} // Student类型标识符，首字母大写包外可见，需要包名，首字母大写标识符，包内可用，不需要包名

// k:=300 短格式不能用在全局定义中
/*
我是一个注释，适合多行注释
*/
// func main() {
// 	// 下面几行的作用
// 	var a int = 100 // 注释，不可以换行
// 	fmt.Println(a)  // TODO:后面完成什么功能
// 	// int := 299 //预定义标识符可以被覆盖，请你一定不要覆盖，你就把它当成关键字好了
// 	// fmt.Println(int)
// 	b := true // 内建的bool类型
// 	fmt.Println(b)
// 	张三 := 100 // 标识符，实际上是个指代
// 	c := "张三" // 这个张三是个值
// 	fmt.Println(c, 张三)
// 	const a1 = 100 // 初始化做了，右边有值，go类型推导
// 	// a1 = 200 不可以，一旦建立就不能改变
// 	c1 := "abc"
// 	fmt.Println(a1, c1)
// 	// const a // 1、未指定类型 2、未初始化，常量必须定义指定关联关系，并不能更改
// 	// const a int // 1、未初始化建立关联关系
// 	const ( // 常量的批量定义
// 		x int    = 100
// 		y string = "xyz"
// 		z        = 1.2
// 	)
// 	// fmt.Printf("%T\n", %x) //常量不能取地址，Go语音的保护
// 	fmt.Println(x, y, z)
// 	// 数组
// 	// const arr =[2]int{1,3} 不可以定义常量，在Go中数组里面的元素是可以发生变化的
// }

// func main() {
// const a = iota
// const b = iota
// fmt.Println(a, b)
// iota用于常量的批量定义
// const (
// 	a = iota // 批量定义,go可以智能补充公式
// 	b        // 上一行继承上一层的公式 b 和 b = iota等价
// 	c
// 	d
// 	_ // blank空白标识符，_可以赋值，但是无法使用其值，黑洞，丢弃类型于Linux的/dev/null
// 	e
// 	f = 567
// 	g // 继承上面的，567
// 	h = iota
// 	i
// )
// fmt.Println(a, b, c, d, e, f, g, h, i)
// const (
// 	a = 100
// 	b
// 	c
// 	d = 123
// 	e
// 	f = iota - 5
// 	g
// )
// fmt.Println(a, b, c, d, e, f, g)

// const (
// 	a = 2 * iota
// 	b
// 	c
// 	d = 2*iota + 1
// 	e
// 	f
// 	g
// )
// fmt.Println(a, b, c, d, e, f)
// 	const a = true
// 	fmt.Println(a)
// }

// func main() {
//		a := true // 标识符指向的值发生变化，这种叫做变量标识符，也叫做变量
//		a = false
//		fmt.Println(a)
// var a // 错误go必须指定类型
// var a string
// var a int // 变量必须指定类型，不然就用短格式做类型推导
// a = 200
// fmt.Println(a)
// var b int // 声明变量，如果不给b值的情况下，默认空值，int默认的值是0
// a := 200
// fmt.Println(a, b)
// fmt.Printf("%T\n", b)
// var c int = 0 // 显式的，明确的
// fmt.Println(c)
// // var d = nil //错误，nil引用类型的零值
// // var a int ,b int  语法错误
// // var a string, b int // 语法错误
// var d, e int // a,b都是int类型，a,b都是零值
// fmt.Println(d, e)
// var a, b int = 100, 200
// fmt.Println(a, b)
// 	var (
// 		a int
// 		b string
// 	)
// 	var (
// 		c int
// 		d = "xyz"
// 	)
// }
// var a int64 = 100 // 左边类型不能省，a int32类型
// fmt.Println(a)
// fmt.Printf("%T\n", a)
// // var a int64 = int64(100) //强制类型转换
// var a // 错误，go必须要给类型
// var a int
// a = 200
// a := 200
// fmt.Println(a)
// var b int // 不仅仅声明，隐式的
// fmt.Printf("%T,%d\n", b, b)
// var c int = 0 //显式的，明确的
// var d = nil //不行，引用类似悬空的值，需要给定类型
// var a,b int //对的，a,b都是int,a,b都被初始化为0
// var a int,b int 语法错误
// var a,b,c string //正确的
// var a int, b string //语法错误
// var a,b int = 1, 5 //对的
// var (
// 	a int
// 	b string

// )
// var (
//
//	c int
//	d = "xyz"
//
// )
// var a int64 = 100 // 左边类型不能省， a int32类型
// fmt.Printf("%T,%d", a, a)
// // var a int64 = int64(100) //强制类型转换
// t := 100
// var a int64 = int64(t)
// fmt.Printf("%T,%d\n", a, a)
// fmt.Printf("%T,%d", t, t)
// var a int
// var b = 100
// a := 1 // 定义赋值+类型推导
//
//		b := 100
//		fmt.Println(a, b)
//		// x,y := 123,"xyz" //可以短格式定义
//		x, y := 123, 345
//		// fmt.Println(x, y)
//		// tmp := y
//		// y = x
//		// x = tmp
//		y, x = x, y
//		// y, _ = x, y
//		fmt.Println(x, y)
//		fmt.Println(g, e, Student{})
//	}
// func main() {
// 	// var a // 错误，无法推测类型
// 	var b int    // 正确，只声明，会自动赋为该类型的零值
// 	var c, d int // 正确，连续的同类型变量可以一并声明，会自动赋为该类型的零值
// 	// var b = 200 // 错误，b多次声明，第二行已经声明过了
// 	fmt.Println(b, c, d)
// }

//	func main() {
//		// 初始化：声明时一并赋初值
//		var a int = 100 // 正确，标准的声明，并初始化
//		b := 200        // 正确，编译根据等式右值推导左边变量的类型
//		// var c = nil // 错误，非法，nil不允许这样用
//		var d, e int = 11, 22 // 正确，var定义多个变量只能是同类型且类型只能写在最后
//		fmt.Println(a, b, d, e)
//	}
// func main() {
// 	// 用var声明，立即赋值，或之后赋值
// 	var b int // 正确，只声明，会自动赋为该类型的零值
// 	b = 200   // 变量可以被重新赋值，要求类型一致
// 	b = 300
// 	// b = "4" // 错误，类型不一致
// 	fmt.Println(b)
// }

// func main() {
// 	// 变量没有定义类型，将使用等号右侧的无类型常量来推断类型
// 	a := 20   // int
// 	b := 3.14 // float64
// 	// 指定变量的类型
// 	// var a int32 = 20
// 	// var b float32 = 3.14
// 	fmt.Println(a, b)
// }

func main() {
	// 批量赋值
	// var a int, b string // 错误，批量不能这么写
	// var ( // 正确，a、b类型知道了，可以用零值
	// 	a int
	// 	b string
	// )
	// var (// 错误，变量必须有类型，但没有给类型，也不能使用值来推导类型
	// 	a
	// 	b
	// )
	// var a int, b string = 111, "abc" // 错误，多种类型不能这么写，语法不对
	// var a, b int
	// fmt.Println(a, b)
	// var (
	// 	a int    = 111
	// 	b string = "abc "
	// ) // 正确，建议批批量常量，变量都这么写
	// fmt.Println(a, b)

	// var (
	// 	a = 111
	// 	b = "abc"
	// )
	// fmt.Println(a, b) // 类型推导
	// 短格式 Short variable declarations
	// _ 空白标识符，或称为匿名变量
	a := 100
	b, c := 200, 300
	// 交换
	b, c = c, b
	fmt.Println(a, b, c)
	d, _, f := func() (int, string, bool) { return 300, "ok", true }()
	fmt.Println(d, f)
}
