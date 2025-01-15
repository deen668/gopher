package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("magedu")
	f, err := os.Open("test")
	if err != nil {
		fmt.Println(err, f)
		return
	}
}
