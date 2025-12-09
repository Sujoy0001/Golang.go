package main

import (
	"fmt"
)

func main() {
	age := 21

	if age >= 18 {
		println("You are an adult.")
	}

	println("//next line")

	if age < 18 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are an adult.")
	}

	println("//next line")

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	println("//next line")

	var role = "admin"
	var hasper = false

	if role == "admin" && hasper {
		fmt.Println("Allow")
	} else {
		fmt.Println("Deny")
	}

	println("//next line")

	if age := 15; age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}
}
