package main

import "fmt"

func names() {

	users := map[string]string{
		"admin": "Steven",
		"user":  "Amina",
	}

	for key, value := range users {
		fmt.Println(key, "=>", value)

	}
}

