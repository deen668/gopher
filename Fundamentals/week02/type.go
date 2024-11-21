package main

import "fmt"

//func main() {
//	//var a = 100
//	//b := 200
//	//var c int = 300
//	//fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
//	//fmt.Printf("%T, %T, %T, %T\n", a, b, c, a+b+c) //format格式化输出
//	var m int64 = 50
//	fmt.Printf("%T,%d\n", m, m)
//	fmt.Printf("%T,%s,%q\n", string(m), string(m), string(m)) //string整数值是ASCII码或Unicode码对应的字符
//	//fmt.Printf("%s\n", 50)  //错误，整数值不能直接转换为字符串
//	fmt.Printf("%v\n", 50)                                         //整数值可以直接输出
//	fmt.Printf("%T,%d\n", rune(m), rune(m))                        //rune类型是int32的别名，可以直接输出,占用4个字节
//	fmt.Printf("%T,%[1]d\n", rune(m))                              //%[1]表示第一个参数，%[2]表示第二个参数
//	fmt.Printf("%[2]T,%[1]d\n", byte(m), string(m))                //byte类型是uint8的别名，可以直接输出,占用1个字节
//	fmt.Printf("%[1]d-- %[3]d; %[2]d++ %d]\n", 100, 200, 300, 400) // %d为什么是300,依次对应,下一个是根据上一个给定的索引进行推导的
//	fmt.Printf("%[1]d-- %d; %[2]d++ %d]\n", 100, 200, 300, 400)
//	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
//	//	//var a uint = -2 //无符号整型不能为负数,溢出了
//	var a = 100
//	//b := 200
//	var b int32 = 200
//	//var c int = 300
//	var c rune = 300 //rune类型是int32类型的别名，4字节
//	fmt.Printf("a=%d %[1]T,b=%d %[2]T,c=%d %[3]T a+b+c=%d %[4]T\n", a, b, c, a+int(b)+int(c))
//	fmt.Println(b + c)
//}

//func main() {
//	type A int32   //定义新类型，本质是一样的，但是定义的数据不能直接通用
//	type B = int32 //类型别名，B和int32是同一个类型的不同名字，B、rune(go内建)、int32都是一个类型的不同名称罢了
//	var b int32 = 200
//	var c rune = 300
//	var d B = 400
//	fmt.Println(b + c + d)
//	var m A = 1000
//	fmt.Println(A(b) + m) //不可以，需要做强制类型转换
//}

//func main() {
//	var a = '测'                       //rune类型 int32 ''表示字符是rune类型，""表示字符串
//	fmt.Printf("%T %[1]d %[1]c\n", a) // 27979 rune类型是int32，是整数，'测'是字面表达，实际上最终存到计算机里面是数字
//	a = 'c'                           //重新赋值,类型不变int32 rune
//	fmt.Printf("%T %[1]d %[1]c\n", a)
//	a = 99                            //是字面常量,默认类型int,a=99是赋值语句不是定义语句,a类型不变
//	a = 0x63                          // 重新赋值，类型不变 int32 rune 0x63是整数标识符
//	fmt.Printf("%T %[1]d %[1]c\n", a) //unicode兼容ascii码,字符c
//	a = '\x63'                        //\x63指16进制字符表示法
//	fmt.Printf("%T %[1]d %[1]c\n", a)
//	a = 27979                         // 有没有隐式类型转换 a = int32(27979)
//	fmt.Printf("%T %[1]d %[1]c\n", a) //unicode码值,字符'测'
//	//b = "测"
//	fmt.Println("====================================")
//	var m int32 = 1000
//	//var n int = int(m) //int32类型强制转换为int类型
//	//n := m
//	var n = 2.6 //float64类型
//	fmt.Printf("%T %d\n", n, int(n))
//	fmt.Println(float64(m) + n) //int32和float64类型不能直接通用,需要强制类型转换
//	fmt.Println(1000 + 2.6)     // untyped constant 1000和2.6是无类型常量,可以直接运算, 特权int往flat64转换
//	fmt.Printf("%T %T %T\n", 1000, 2.6, 1000+2.6)
//	fmt.Printf("%f\n", 1000+2.468)             //%f输出浮点数.默认精度6位
//	fmt.Printf("%.2f\n", 1000+2.468)           //%.2f输出浮点数,保留2位小数
//	fmt.Printf("%30f\n", 1000+2.468)           //%30f输出浮点数,宽度为30,默认精度6位,默认右对齐
//	fmt.Printf("%7.1f\n", 1000+2.468)          //%7f输出浮点数,宽度为7,默认精度1位,右对齐,保留小数点1位
//	fmt.Printf("%-30f========\n", 1000+2.468)  //%-30f输出浮点数,宽度为30,默认精度6位,左对齐
//	fmt.Printf("%-7.1f========\n", 1000+2.468) //%-7.1f输出浮点数,宽度为7,默认精度1位,左对齐,保留小数点1位
//
//}

func main() {
	//var m int = '测'
	//fmt.Println(m)
	//var m byte = '测' //byte类型,1个字节,放不下'测',会报错,uint8 int32==>27979 unicode不是utf-8
	//var b = "测"                       //string类型,utf-8编码,3个字节
	var m byte = 'a'                  //byte类型,1个字节,放得下'a' var m byte = byte('a') 0x61 97
	fmt.Printf("%T %[1]d %[1]c\n", m) // ascii码值,字符'a',不超过一个字节
}
