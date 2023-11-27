package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	N := 0
	for i := B - 1; i > A; i-- {
		fmt.Println(i)
		N++
	}
	fmt.Printf("Количество чисел: %d\n", N)
}
