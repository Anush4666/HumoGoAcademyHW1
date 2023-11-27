package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min_ := arr[0]
	max_ := arr[n-1]

	for i := 0; i < n; i++ {
		if min_ > arr[i] {
			min_ = arr[i]
		}
		if max_ < arr[i] {
			max_ = arr[i]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] == max_ {
			arr[i] = min_
		}
	}

	fmt.Println(arr)
}
