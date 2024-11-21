package main

import "fmt"

//导入包,第三方包

//var ( //全局变量，顶层代码
//	a = 100
//	b = "abc"
//)

//func main() {
//	var a = 100 //右边是无类型常量,缺省类型是int,左边a没有指定类型，用到类型推导，所以a是int类型
//	a = 200
//	a = 123
//	fmt.Println(a)
//}

//func main() {
//    //var a // 错误的，因为没有右边,go是强类型语言,此时不能利用右边来推导
//    //var a int // 正确的,因为指定了类型，但是没有赋值，在go中正常的原因是go中有默认值，零值可用，只用变量可以不赋值
//    //const b int  // 错误的，因为常量必须赋值
//    //a = 100
//    var a int32 = 200 //正确,指定了类型，也赋值了，不会采用类型推导
//    fmt.Println(a)
//    a = 100
//    //var a int = 100 //不行，重复定义
//    //var c = nil //nil是预定义的标识符，表示指针、通道、函数、接口、映射或切片类型的零值，不能用于普通类型
//
//}

//func main() {
//	//var (
//	//	a = 100
//	//	b int
//	//	c string = "abc"
//	//	//d 错误的，无法推导，用不上零值
//	//)
//	//fmt.Println(a, b, c)
//	fmt.Println(a, b)
//	//var a int b int //错误的，不能连续定义
//	//var a int, b int //错误的，不能连续定义
//	var a1, b1 int //如果要写到一行,必须同类型，只需要在最后指定类型
//	a1, b1 = 100, 200
//	fmt.Println(a1, b1)
//	var c1, d1 int = 100, 200 //如果要写到一行,必须同类型，只需要在最后指定类型,可以赋值，但是需要全部给出
//	fmt.Println(c1, d1)
//}

// func main() {
// a := "mn" //类型推导，短格式变量定义语句,定义变量,定义了变量标识符，右边可以用来推导a的类型为int
// //b        //不是短格式,什么都不是,也没有var，const
// b := "abc"
// fmt.Println(a, b)
// b = "a"
// b = a //变量交换
// fmt.Println(a, b)
// var a := 100 //错误
// a int :=200 //错误
// var (
//
//	a1 := 100  //错误
//
// )
// fmt.Println(a1)
//
//		a, b, c := 100, 200, 300 //这里a不能重新定义为新类型，a被检测到了，go语音上只能迁就你了，但是要求a同类型，b是新定义的,左右必须一一对应
//		fmt.Println(a, b, c)
//		a, b = b, c
//		fmt.Println(a, b, c)
//		var (
//			m = 111
//			n = 222
//		)
//		m, n = n, m //交换,赋值语句=,右边先做，照相定格，左边再做
//		tmp := m
//		m = n
//		n = m
//		fmt.Println(m, n, tmp)
//	}
//const a = 100
//
//var b = 200
//
//func main() { //main函数叫做入口函数，go约定main函数必须在main包中定义
//	a, _, c := 123, 345, 678 //_blank空白标识符，只能用在左边,其他地方不可用
//	fmt.Println(a, c)        //编译了，生成可执行文件并执行，内存中有个进程，进程内的内存空间有了你的数据
//	//fmt.Println(d)
//	fmt.Println(E)
//}

func main() {
	var a = 100
	var b int32 = 200
	fmt.Println(a, b)
	c := true
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
