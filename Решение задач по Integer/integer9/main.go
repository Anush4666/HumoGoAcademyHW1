package main

import "fmt"

func main() {
	var abc int
	fmt.Println("Ввести только трехзначное число!")
	fmt.Scan(&abc)
	fmt.Println(abc / 100)

}
