package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} 

	// infinite loop
	// for {
	// 	println(i)
	// }
	// traditional for loop

	for i:= 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// clssic for loop
	count := 10
	for i := 1; i < count; i++ {
		// break
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i := range 10 {
		println(i)
	}
}