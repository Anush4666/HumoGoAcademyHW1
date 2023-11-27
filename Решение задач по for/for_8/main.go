package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c int = 1
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		c *= i
	}
	fmt.Println(c)
}
