//package main
//
//import "fmt"
//
//func main() {
//	var N int
//	fmt.Scan(&N)
//	array := make([]int, N)
//	for i := 0; i < N; i++ {
//		fmt.Scan(&array[i])
//	}
//	for i := N - 1; i >= 0; i-- {
//		fmt.Printf("%v ", array[i])
//	}
//}



//Второй способ

//package main
//
//import "fmt"
//
//func main() {
//	var N int
//	fmt.Scan(&N)
//	array := make([]int, N)
//	for i := 0; i < N; i++ {
//		fmt.Scan(&array[i])
//	}
//	for i := 0; i < N/2; i++ {
//		array[i], array[N-1-i] = array[N-1-i], array[i]
//
//	}
//	fmt.Println(array)
//}
