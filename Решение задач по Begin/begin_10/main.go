package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println((a * a) + (b * b))
	fmt.Println((a * a) - (b * b))
	fmt.Println((a * a) * (b * b))
	fmt.Println((a * a) / (b * b))
}
