package main

import "fmt"

func main() {
	var a, b, c int
	var num = 0
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		num++
	}
	if b > 0 {
		num++
	}
	if c > 0 {
		num++
	}
	fmt.Println("Количество положительных чисел", num)
}
