package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used data structure in go

func main1() {
	// uninitialized slice is nil
	// var nums []int 
	// fmt.Println(nums)

	// var nums = make([]int, 0, 5) // length, capacity
	// fmt.Println(nums)
	// fmt.Println("Length:", len(nums))
	// fmt.Println("Capacity:", cap(nums))
	// fmt.Println(nums)

	// nums = append(nums, 1)
	// fmt.Println(nums)
	// nums = append(nums, 2, 3, 4, 5, 6)
	// fmt.Println(nums)
	// fmt.Println("Length:", len(nums))
	// fmt.Println("Capacity:", cap(nums))

	// nums = append(nums, 1)
	// fmt.Println(nums)
	// nums = append(nums, 2, 3, 4)
	// nums = append(nums, 2, 3, 4)
	// fmt.Println(nums)
	// fmt.Println("Length:", len(nums))
	// fmt.Println("Capacity:", cap(nums))

	var nums = []int{1,2,3,4,5}
	fmt.Println(nums[0:3])

	var num1 = []int{1,2,3,}
	var num2 = []int{1,2,3}

	fmt.Println(slices.Equal(num1, num2))
}