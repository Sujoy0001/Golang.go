package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 3, 4, 5, 8, 4, 3, 2, 5, 7}
	nums1 := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				nums1 = append(nums1, i)
			}
		}
	}

	fmt.Println(nums)
}