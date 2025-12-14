package main

import "fmt"

func main() {

	// sum of array elements
	
	// array := [5]int{1, 2, 3, 4, 5}

	// sum := 0

	// for i := range len(array) {
	// 	sum += array[i]
	// }  

	// fmt.Println(sum)

	// max number in array
	// array := [5]int{70,58,99,41,85}

	// max := 0

	// for i := range len(array) {
	// 	if array[i] > max {
	// 		max = array[i]
	// 	}
	// }

	// fmt.Println(max)

	// count even number
	array := [5]int{1,5,8,6,4}

	count := 0

	for i := range len(array) {
		if array[i] % 2 == 0 {
			count += 1
		}
	}

	fmt.Println(count)


}