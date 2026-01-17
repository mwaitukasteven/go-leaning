package main

import "fmt"

func main() {

	greetUser("Steven")
	greetUser("Amina")

}

func greetUser(name string) {
	fmt.Println("Hello", name)
}
