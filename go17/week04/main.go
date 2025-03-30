package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// a := 10 // 只能定义函数内部的变量
	// if a < 0 {
	// 	fmt.Println("a is regative")
	// } else if a > 0 { // 能走到这里，说明已经测试过了a<0,也就是这里隐含一个条件a>=0
	// 	fmt.Println("a is positive")
	// } else { // 能走到这里说明已经测试通过了a<0和a>0，隐含这一个条件a==0
	// 	fmt.Println("is zero")
	// }
	// if a < 0 {
	// 	fmt.Println("a is regative")
	// } else {
	// 	if a > 0 {
	// 		fmt.Println("a is positive")
	// 	} else {
	// 		fmt.Println("is zero")
	// 	}
	// }

	// switch a := 10; a { //:=声明，赋值
	// case 0:
	// 	fmt.Println("zero")
	// case 10:
	// 	fmt.Println("10 go")
	// 	// fallthrough // 一般不要加
	// case 100:
	// 	fmt.Println("100")
	// case 40:
	// 	fmt.Println("40")
	// 	fmt.Println(a)
	// }
	// fmt.Println(a)
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// switch { // 不写代表写了bool值true
	// case a > 0:
	// 	fmt.Println("a is regative")
	// case a < 0:
	// 	fmt.Println("a is positive")
	// default:
	// 	fmt.Println("a zero")
	// }

	// switch a:=10; {
	// case a >0:
	// 	fmt.Println("a is regative")
	// case a <0 && a ==0:
	// 	fmt.Println("negative")
	// case a ==0:
	// default:
	// 	fmt.Println("zero")
	// }

	// if score,line :=90,90; score >line{
	// 	fmt.Println("good")
	// } else{
	// 	fmt.Println("bod",score)
	// }
	// k := 0                       // var k =0
	// for k := 100; k < 105; k++ { // 作用域的问题，for后面的:重新定义只能在for中使用的一个k
	// 	fmt.Println(k, "~~~~~")
	// }
	// fmt.Println(k)

	// k := 0                       // var k =0
	// for k := 100; k < 105; k++ { // 作用域的问题，for后面的:重新定义只能在for中使用的一个k
	// 	go func() {}()          // go func (){}这是一个函数定义，函数后面再补一个()这是函数调用这是gorutime
	// 	// time.Sleep(time.Second) // 睡1s

	// 	fmt.Println(k)
	// }
	// fmt.Println(k)

	// for k := 100; k < 105; k++ {
	// 	go func() {
	// 		time.Sleep(1 * time.Second) // 睡1s就这样写，2秒就*2
	// 		fmt.Println(k)
	// 	}()
	// }
	// time.Sleep(2 * time.Second)
	// for k := 100; ; /*k < 105*/ {  //当true看
	// 	fmt.Println(k)
	// 	time.Sleep(time.Second)
	// }
	// k := 100
	// for {
	// 	fmt.Println(k)
	// }
	// for k = 5; k < 10; k-- {
	// 	fmt.Println(k)
	// 	time.Sleep(time.Second)
	// }
	// time.Sleep(2 * time.Second)
	// for j, k := 5, 6; ; {
	// 	fmt.Println(j, k)
	// 	j++
	// 	k++
	// 	time.Sleep(time.Second)
	// }
	// time.Sleep(2 * time.Second)
	// for i := 1; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	if i&1 == 1 {
	// 		fmt.Println(i)
	// 	}
	// }
	// for i := 1; i < 10; i += 2 {
	// 	fmt.Println(i)
	// }
	// i := -1
	// for i = 0; ; i++ {
	// 	fmt.Println(i)
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	// break // 终止for循环
	// }
	// fmt.Println(i, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// LA: //label
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		fmt.Println(i, j, i*j)
	// 		if j == 2 {
	// 			goto END
	// 			// break LA // break j loop
	// 		}
	// 		fmt.Println(i, j, i*j)
	// 	}
	// END:
	// 	fmt.Println("over~~~~~~")

	// 	// if i == 2 {
	// 	// 	break //break j loop
	// 	// }
	// }
	// a := [3]int{5, 7, 9} //[]int切片，[]int{1,2,3}数组,{}可以写元素
	// for i, v := range a {
	// 	fmt.Println(i, v)
	// }
	// for i := range a {
	// 	fmt.Println(i)
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// m := map[string]int{
	// 	"abc": 100,
	// 	"xyz": 200,
	// 	"cde": 300,
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for k := range m {
	// 	fmt.Println(k)
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// m := [3]int{5, 7, 9}
	// for i := 0; i < len(m); i++ {
	// 	fmt.Println(m[i])
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for i, v := range m {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// for i := range 5 { //range n[0,n)
	// 	fmt.Println(i)
	// }
	// a := [3]int{5, 7, 9}
	// for i := range 3 { // range n[0,n)
	// 	fmt.Println(i, a[i])
	// }
	// s := "abcd测试"
	// for i, v := range s {
	// 	// fmt.Println(i, v)
	// 	fmt.Printf("%d:%T %[2]v %[2]q\n", i, v) // 遍历按照字符
	// }
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%d:%T %[2]v %[2]q\n", i, s[i]) // 遍历按照字节
	// }
	// os.Setenv("GODEBUG", "randautoseed=0")
	// 随机数
	for i := range 10 { // 1.20版本随机数种子
		fmt.Println(i, rand.Intn(10)) //[0,9)
	}
	// rand.Seed(time.Now().UnixNano())
	// for i := range 10 {
	// 	fmt.Println(i, rand.Intn(10)) //[0,9)
	// }
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	seed := time.Now().UnixNano() // 获取当前时间戳
	src := rand.NewSource(seed)   // 创建新的随机数种子
	r := rand.New(src)            // 创建新的随机数生成器
	for i := range 10 {
		fmt.Println(i, r.Intn(10)) //[0,9)
	}
}
