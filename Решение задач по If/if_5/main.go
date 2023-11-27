package main

import "fmt"

func main() {
	var a, b, c int
	var num = 0
	var num1 = 0
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		num++
	} else if a < 0 {
		num1++
	}
	if b > 0 {
		num++
	} else if b < 0 {
		num1++
	}
	if c > 0 {
		num++
	} else if c < 0 {
		num1++

	}
	fmt.Println("Количество положительных чисел", num)
	fmt.Println("Количество отрицательных чисел", num1)
}
