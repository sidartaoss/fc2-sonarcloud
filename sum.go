package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
	// fmt.Println(sumX(10, 10))
	// fmt.Println(sub(10, 10))
	// fmt.Println(multiplication(10, 10))
	// fmt.Println(division(10, 10))
}

func sum(a, b int) int {
	return a + b
}

// func sumX(a, b int) int {
// 	return a + b + a
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func multiplication(a, b int) int {
// 	return a * b
// }

// func division(a, b int) int {
// 	return a / b
// }
