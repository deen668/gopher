package main

import "fmt"

// 同一个目录就是同一个包，同一个目录只能有一个包
var d = "abc" //顶层代码，全局变量，包内可见，可以用，包外不可见
var E = "xyz" //首字母大写，可以被其他包访问
//比如说这里有函数，结构体等要用这个E

func main() {
	//var a int8 = -2
	//fmt.Println(a)
	var a = 100
	var b int64 = 200
	var c rune = 300 //rune类型是int32的别名，可以直接输出,占用4个字节
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
