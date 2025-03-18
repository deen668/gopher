package main

import "fmt"

func main() {
	a := 100
	// if a < 0 {
	// 	fmt.Println("a is regative")
	// } else if a > 0 { // 能走到这里，说明已经测试过了a<0,也就是这里隐含一个条件a>=0
	// 	fmt.Println("a is positive")
	// } else { // 能走到这里说明已经测试通过了a<0和a>0，隐含这一个条件a==0
	// 	fmt.Println("is zero")
	// }
	if a < 0 {
		fmt.Println("a is regative")
	} else {
		if a > 0 {
			fmt.Println("a is positive")
		} else {
			fmt.Println("is zero")
		}
	}
}
