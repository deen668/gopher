package main

import "fmt"

func main() {
	var a = 100 // a是标识符,var variable 变量
	a = 200
	fmt.Println(a)
	const b = 300
	//b = 400 // Error: cannot assign to b
	fmt.Println(b)
	const c = "abc" //类型推导
	fmt.Println(c)
	const d = 100       //字面常量也叫无类型常量关联，100和d关联，默认类型int
	const e uint8 = 100 //右边无类型常量100和左右有类型常量e关联  int 8bit ==> 1byte 0--127 0--255 int和uint是有符号和无符号的区别
	const f = 1.5
	fmt.Println(d, e, f)
	const g = true
	fmt.Println(g)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}
