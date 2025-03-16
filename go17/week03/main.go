package main

import "fmt"

func main() {
	// var a int8 = 127 // int8(127)--->做强制类型转换，128不行，溢出了
	// fmt.Println(a)
	// a = a + 1 // go可以做int8的值+1但是比较危险，不建议做
	// fmt.Println(a)
	// fmt.Println(1 + 2)
	// fmt.Println(1 + 3.1) // 字面常量做字面常量隐式转换
	// fmt.Printf("%T\n", 1+2+3.1)
	// a := 100
	// b := 100
	// var c int = -600
	// fmt.Printf("%T,%T,%T\n", a, b, a+b+c)
	// fmt.Printf("%+v,%+v,%+v\n", a, b, a+b+c)
	// fmt.Printf("%d,%d,%d\n", a, b, a+b+c)
	// var d byte = 0x10
	// fmt.Println(a + int(d))
	// fmt.Println(byte(c))
	// f := 1.2
	// fmt.Println(int(f))
	// fmt.Println(byte(f))
	// fmt.Printf("%T,%d,%T,%v\n", a, a, float64(a), float64(a))
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 浮点数
	// f := math.Pi
	// fmt.Printf("1 %f\n", f)        //%f, 精度为6，保存小数点后6位
	// fmt.Printf("2 %3.f\n", f)      // %f, 精度为3，保存小数点后3位
	// fmt.Printf("3 %3f\n", f)       //%f, 精度为3，宽度3，显示数据超过宽度，撑爆了
	// fmt.Printf("4 |%10f|\n", f)    // 默认精度为6，前面10表示显示宽度10
	// fmt.Printf("5 |%10.3f|\n", f)  // 精度3，宽度10
	// fmt.Printf("6 |%-10.3f|\n", f) // 精度3，宽度10, -号表示左对齐
	// fmt.Printf("7 |%-.3f|\n", f)
	// f = 0.00001234567979
	// fmt.Printf("8 |%-.3e|\n", f) // %e科学计数法
	// fmt.Printf("8 |%-.3g|\n", f) // %g，智能选择%g或者%f
	// var a = 1 // a int
	// var b = 3.1 // b bloat64
	// a + b
	// fmt.Println(a+ b) //右边的无类型字面常量的缺省类型就是标识符的类型,a和b标识符不同类型，不能相加
	// fmt.Println('\rr\nn') //rune类型int32，单字节，超界了
	// 	fmt.Println("1 \rr\nn")
	// 	fmt.Println(`2 ab	\tc`)
	// 	fmt.Println("3 ab`c")
	// 	fmt.Println("4 ab\n\tc")
	// 	s1 := `ab
	// c
	// 	d`
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Printf("%d\n", len(s1))
	// 	fmt.Println(s1)
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 	s2 := "ab\nc\n\td"
	// 	fmt.Println(s2)
	// 	fmt.Printf("%d\n", len(s2))
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("abc" + `xyz`)
	// fmt.Println('\x61') // 1个字符，4个字节有符号正整数
	// fmt.Printf("%d\n", '\x61')
	// fmt.Println("\x61")
	// fmt.Printf("%s\n", "\x61")
	// a := 0x61         // 97 int 8 bytes,7个高字节都是0,最低字节是0b0110 0001-->0x61
	// b := '\x61'       // rune int32 4 bytes，3个高字节都是0，最低字节是0b0110 0001
	// var c byte = 0x61 // byte uint8 1 bytes，无符号数，0b0110 0001
	// d := "\x61"       // string 标头值+底层字节数组，字节数组有1个元素
	// fmt.Printf("%T,%T, %T, %T\n", a, b, c, d)
	// fmt.Printf("%T %[1]d,%T %[2]d %[2]x %#[2]x %#[2]b, %T %[3]d %#[3]o, %T %[4]s\n", a, b, c, d)
	// fmt.Println(len(d))
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// fmt.Printf("%T %[1]d %[1]x %[1]b %[1]c %[1]q\n", a)
	// fmt.Printf("%T %[1]d %[1]x %[1]b\n", b)
	// fmt.Printf("%T %[1]s %[1]q\n", d)
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// s1 := fmt.Sprintf("%d,%d,%d------%s", a, b, c, d)
	// fmt.Println(s1)
	// fmt.Println(s1 + "\tabc")
	// fmt.Println("abcd123\t" + s1)
	// fmt.Println(5/2, 5/-2) // 整除，截取整数部分，整数/整数=整数，不能是浮点数
	// fmt.Println(5 / 2.0)   // 隐式类型转换，只有无类型字面常量可以做隐式类型转换
	// a := 5
	// b := 2.5
	// fmt.Println(float64(a) / b) // 这种不行，定位完类型后不能做类型转换go是强类型语言
	// a := 0xe
	// fmt.Println(a&5, a&^5, a&1)
	// fmt.Println(2|1, 3|1)
	// fmt.Println(2^1, 2^2)
	// fmt.Println(2>>1, 2/2, 8>>1, 8>>2, 8>>3, 8>>4)
	// fmt.Println(8/2, 8/4, 8/8, 8/16)
	// fmt.Println(1<<1, 1^2, 1<<2, 1*4)
	// fmt.Println(1 == 'a')
	// // fmt.Println(1 == "a")
	// fmt.Println(1 == 2.1)
	// var a =1
	// var b = 2.3
	// fmt.Println(a == b)
	// "abc" && "xyz" //&& 逻辑运算符不能连接除bool值以外的值
	// a := 100
	// fmt.Println(a > 5 && "abc" > "a" || false)
	// fmt.Println("abc" > "a")
	a := 100
	b := &a //*int是我是一个指针，b取得是内存地址
	fmt.Printf("%T,%[1]p\n", b)
	fmt.Println(*b)
	c := b
	fmt.Println(c)
	fmt.Println(a == *b, *b == *c)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	t := 100
	fmt.Println(a == t, b == &t)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var p1 *int // 零值nil，悬空指针，非常危险
	fmt.Println(p1)
	p1 = &a
	fmt.Println(p1)
	*p1 = 333
	fmt.Println(a, *b, *c, *p1)
}
