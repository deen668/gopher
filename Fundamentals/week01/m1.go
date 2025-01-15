package main

import "fmt"

func main() {
	const (
		a = iota + 1
		b // 批量定义的时候，下面一行可以继承上面的表达式，b = iota
		c
		_ // 空白标识符，用于占位
		d
		t = 100
		_
		_
		e
	)
	fmt.Println(a, b, c, d, t, e)
}
