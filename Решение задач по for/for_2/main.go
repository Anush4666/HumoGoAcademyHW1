package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	N := 0
	for i := A; i <= B; i++ {
		fmt.Println(i)
		N++
	}
	fmt.Printf("Количество чисел: %d\n", N)
}
