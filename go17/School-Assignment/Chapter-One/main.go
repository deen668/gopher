// 打印九九乘法表。如果可以要求间隔均匀
// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 9; i++ { // 外层循环，控制行数
// 		for j := 1; j <= i; j++ { // 内层循环，控制列数
// 			fmt.Printf("%d*%d=%d\t", i, j, i*j) // 打印乘法表
// 		}
// 		fmt.Println() // 换行
// 	}
// }

// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 9; i++ {
//			for j := 1; j <= i; j++ {
//				// 使用%-2d确保数字左对齐，占用2个字符位置
//				fmt.Printf("%d×%d=%-2d  ", j, i, i*j)
//			}
//			fmt.Println()
//		}
//	}
//
// 2、随机生成20以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个（不是索引）累加求积，最后打印出结果
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	// "time"
// )

// func main() {
// 	// rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
// 	num := make([]int, 20) // 创建一个长度为20的整数切片

// 	for i := 0; i < 20; i++ { // 生成20个非0正整数
// 		num[i] = rand.Intn(99) + 1 // 生成一个1到99之间的随机数
// 	}

// 	fmt.Println("生成的20个非0正整数：", num) // 打印生成的20个非0正整数
// 	sum := 0                         // 初始化sum
// 	product := 1                     // 初始化product

// 	for i := 0; i < 20; i++ { // 遍历20个非0正整数
// 		if i%2 == 0 { // 判断是否为偶数
// 			sum += num[i] // 累加偶数
// 		} else {
// 			product *= num[i] // 累乘奇数
// 		}
// 	}

// 	fmt.Printf("单数个数的和：%d\n", sum)
// 	fmt.Printf("偶数个数的积：%d\n", product)
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	// r := rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机数种子,1.20版本一下需要
// 	var (
// 		sum            = 0 // 初始化sum
// 		product uint64 = 1 // 如果乘积可能超过int上界，就使用uint64
// 	)
// 	for i := 0; i < 20; i++ { // 生成20个非0正整数
// 		v := rand.Intn(99) + 1 // 生成一个1到99之间的随机数
// 		fmt.Println(v)         // 打印随机数
// 		if i%2 == 0 {          // 判断是否为偶数
// 			sum += v // 累加偶数
// 		} else {
// 			product *= uint64(v) // 累乘奇数
// 		}
// 	}
// 	fmt.Printf("单数个数的和：%d\n", sum)
// 	fmt.Printf("偶数个数的积：%d\n", product)
// }

// 3、打印100以内的斐波那契数列
// package main

// import "fmt"

//	func main() {
//		a, b := 1, 1      // 初始化a和b
//		limit := 100      // 设置上限
//		sum := 0          // 初始化sum
//		product := 1      // 初始化product
//		fmt.Println(a, b) // 打印前两个数
//		for {
//			a, b = b, a+b   // 更新a和b
//			if a >= limit { // 如果a大于等于上限，则退出循环
//				break
//			}
//			fmt.Println(a) // 打印a
//			sum += a
//			product *= a
//		}
//		fmt.Println(sum)     // 打印100以内的斐波那契数列的和
//		fmt.Println(product) // 打印100以内的斐波那契数列的积
//	}
//
// 4、打印100以内的所有质数,并计算100以内的所有质数的和和积
package main

import "fmt"

func main() {
	sum := 0             // 初始化sum
	product := uint64(1) // 如果乘积可能超过int上界，就使用uint64
	for i := 2; i <= 100; i++ {
		isPrime := true // 初始化isPrime,默认是质数
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false // 如果i能被j整除，则i不是质数
				break
			}
		}
		if isPrime {
			fmt.Println(i)       // 打印质数
			sum += i             // 累加质数
			product *= uint64(i) // 累乘质数
		}
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(sum) // 打印100以内的所有质数的和
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(product) // 打印100以内的所有质数的积
}
