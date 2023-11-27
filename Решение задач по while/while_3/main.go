package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Print("k: ")
	fmt.Scan(&k)

	count := 0
	for n >= k {
		n -= k
		count++
	}

	fmt.Println("Частное от деления нацело N на K =", count, "Остаток от этого деления =", n)
}
