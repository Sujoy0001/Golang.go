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
	for j := 1; j <= 3; j++ {
		fmt.Println(i * j)
	}
}