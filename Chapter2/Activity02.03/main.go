package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 20)
	for i := range nums {
		nums[i] = rand.Intn(10)
	}
	fmt.Println("Before:", nums)
	for l := len(nums) - 1; l > 0; l-- {
		for i := 0; i < l; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	fmt.Println("After:", nums)
}
