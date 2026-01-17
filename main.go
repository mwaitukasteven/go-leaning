package main

import "fmt"

func main() {

	fmt.Println("Inside main function")
	sayHello()
	showMessage()

}

func sayHello() {
	fmt.Println("Hello from sayHello function")
}

func showMessage() {
	fmt.Println("Go is easy to learn")
}
