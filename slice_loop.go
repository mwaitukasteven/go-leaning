package main

import "fmt"

func main() {

	names := []string{"Steven", "Amina", "John"}

	for index, name := range names {
		fmt.Println(index, name)
	}
}
