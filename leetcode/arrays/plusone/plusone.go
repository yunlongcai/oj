package main

import "fmt"

func main() {
	numss := [][]int{
		{
			0,
		},
		{
			9,
		},
		{
			1, 2, 3,
		},
		{
			4, 3, 2, 1,
		},
		{
			9, 9, 9, 9,
		},
		{
			1, 2, 9, 9,
		},
	}

	for _, nums := range numss {
		fmt.Println(plusOne(nums))
	}
}

// 改进版，出现没有进位的，就直接返回了。
func plusOne(digits []int) []int {
	var carry bool
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + 1
		digits[i] = v % 10
		if carry = v/10 == 1; !carry {
			return digits
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 原始解题思路，每次判断是否有进位，遍历完所有的
func plusOne1(digits []int) []int {
	var carry bool
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i]
		if carry || i == len(digits)-1 {
			v++
		}
		digits[i] = v % 10
		carry = v/10 == 1
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
