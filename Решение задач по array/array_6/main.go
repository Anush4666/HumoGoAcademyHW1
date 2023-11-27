package main

import "fmt"

func main() {
	var (
		n       int
		counter int
		x       int
	)
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&x)

	for i := 0; i < n; i++ {
		if arr[i] == x {
			counter++
		}
	}

	fmt.Println(counter)
}
