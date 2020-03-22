package main

import "fmt"

func main() {
	numss := [][]int{
		{
			2, 2, 1,
		},
		{
			4, 1, 2, 1, 2,
		},
	}
	for _, nums := range numss {
		fmt.Println(singleNumber(nums))
	}
}

func singleNumber(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		n ^= nums[i]
	}
	return n
}
