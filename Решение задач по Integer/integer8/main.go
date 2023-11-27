package main

import "fmt"

func main() {
	var ab int
	fmt.Println("Ввести только двухзначное число!")
	fmt.Scan(&ab)
	fmt.Println((ab / 10) + (ab%10)*10)

}
