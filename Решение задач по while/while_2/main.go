package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)

	count := 0
	for a >= b {
		a -= b
		count++
	}

	fmt.Println("Количество отрезков Б: ", count)
}
