package main

import "fmt"

func main() {
	// TODO: 我有某个功能待实现
	fmt.Println("magedu")
	fmt.Println(".com")
	//var bool = true
	//fmt.Println(bool)
	var a = 100
	fmt.Println(a) //定义个a标识符，variable 变量
	a = 200
	fmt.Println(a)
	const b = 200
	//b = 400 //常量不能被修改
	fmt.Println(b)
	const (
		c       = "abc" //类型推导
		d int   = 100   //字面常量无类型常量100关联
		e uint8 = 100   //无类型常量100和左边有类型常量e uint8关联
		// int 8bits ==> 1byte 可以表示的有符号数值范围是-128~127，整数部分0~127，负数部分-128~-1，0~255
		f float32 = 1.5
		g         = true
	)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("%T\n", c) // %T 类型
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}
