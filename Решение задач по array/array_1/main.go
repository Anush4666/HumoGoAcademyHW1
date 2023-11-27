package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	for i := 1; i < N; i++ {
		if (array[i] > 0 && array[i+1] > 0) || (array[i+1] < 0 && array[i] < 0) {
			fmt.Println("Yes")
			break

		} else {
			fmt.Println("No")
			break
		}
	}
}
