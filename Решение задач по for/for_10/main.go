package main

import (
	"fmt"
)

func main() {
	var (
		a   int
		sum float64
	)
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		sum += float64(1/i)
	}
	fmt.Println(sum)
}
