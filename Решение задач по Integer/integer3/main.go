package main

import "fmt"

func main() {
	var B int
	fmt.Println("Введите четырёхзначное число и выше!")
	fmt.Scan(&B)
	fmt.Println(B / 1024)

}
