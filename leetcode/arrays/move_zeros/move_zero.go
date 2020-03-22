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
		{
			0, 1, 0, 3, 12,
		},
		{
			0, 0, 0, 0, 1, 0, 3, 12, 0, 0, 0,
		},
	}
	for _, nums := range numss {
		moveZeroes(nums)
		fmt.Println(nums)
	}
}

func moveZeroes(nums []int) {
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != 0 {
	// 		continue
	// 	}
	// 	for j := i; j < len(nums); j++ {
	// 		if nums[j] == 0 {
	// 			// 与首个不是0 的交换
	// 			for k := j + 1; k < len(nums); k++ {
	// 				if nums[k] != 0 {
	// 					nums[j], nums[k] = nums[k], nums[j]
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	//cur 表示下个放置非 0 数字的位置
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if cur != i {
				nums[cur], nums[i] = nums[i], nums[cur]
			}
			cur++
		}
	}
}
