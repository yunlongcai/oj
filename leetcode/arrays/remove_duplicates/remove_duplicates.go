package main

import "fmt"

func main() {
	testcases := []struct {
		nums []int
	}{
		{
			nums: []int{
				1, 2, 3,
			},
		},
		{
			nums: []int{
				1, 1, 2,
			},
		},
		{
			nums: []int{
				0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
			},
		},
	}
	for _, tc := range testcases {
		n := removeDuplicates(tc.nums)
		fmt.Println(tc.nums[:n])
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curVal := nums[0]
	curPos := 1
	for i := 1; i < len(nums); i++ {
		if curVal == nums[i] {
			continue
		}
		nums[curPos] = nums[i]
		curVal = nums[i]
		curPos++
	}
	return curPos
}
