package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch

	i := 7

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3: 
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	}

	// multiple cases switch

	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday:
		fmt.Println("It's the start of the week.")
	case time.Wednesday, time.Thursday:
		fmt.Println("It's midweek.")
	case time.Friday:
		fmt.Println("It's almost the weekend!")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	}

	// type switch

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's a integer value")
		case string:
			fmt.Println("it's a string value")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println(t)
		}
	}

	whoami(true)
}