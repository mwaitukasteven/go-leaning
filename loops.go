package main

import "fmt"

func look() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}

}


func mai() {

	i := 1

	for {
		fmt.Println("Looping:", i)
		i++

		if i > 5 {
			break
		}
	}

}


func ma() {

	count := 1

	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

}