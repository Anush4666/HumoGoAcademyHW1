package main

import (
	"fmt"
)

func main() {
	var (
		n    int
		x    int
		flag bool
	)
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&x)

	for i := 0; i < n; i++ {
		if arr[i] == x {
			flag = true
			break
		}
	}

	if flag {
		fmt.Println(x)
	} else {
		fmt.Println(x - 1)
	}

}
