package main

import (
	"fmt"
)

func main() {
	var K, N int
	fmt.Scan(&K, &N)

	for i := 0; i < N; i++ {
		fmt.Println(K)
	}
}
