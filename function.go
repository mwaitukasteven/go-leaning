func add(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
