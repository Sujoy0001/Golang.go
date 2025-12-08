package main

import "fmt"
const age = 21
var nameIS = true

const (
	port = 5000
	host = "localhost"
)

func main() {
	// var name string = "Sujoy"

	var name = "Sujoy"
	var isAdult = true

	myname := "Sujoy garai"

	var number int

	number = 27

	fmt.Println(name)
	fmt.Println("Is Adult:", isAdult)
	fmt.Println("My Name is:", myname)
	fmt.Println("Number:", number)
	fmt.Println(age)
	fmt.Println(nameIS)
	fmt.Println(port, host)
}