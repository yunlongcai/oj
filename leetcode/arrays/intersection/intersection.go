package main

import "fmt"

func main() {
	numss := [][2][]int{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
		},
		{
			[]int{54, 93, 21, 73, 84, 60, 18, 62, 59, 89, 83, 89, 25, 39, 41, 55, 78, 27, 65, 82, 94, 61, 12, 38, 76, 5, 35, 6, 51, 48, 61, 0, 47, 60, 84, 9, 13, 28, 38, 21, 55, 37, 4, 67, 64, 86, 45, 33, 41},
			[]int{17, 17, 87, 98, 18, 53, 2, 69, 74, 73, 20, 85, 59, 89, 84, 91, 84, 34, 44, 48, 20, 42, 68, 84, 8, 54, 66, 62, 69, 52, 67, 27, 87, 49, 92, 14, 92, 53, 22, 90, 60, 14, 8, 71, 0, 61, 94, 1, 22, 84, 10, 55, 55, 60, 98, 76, 27, 35, 84, 28, 4, 2, 9, 44, 86, 12, 17, 89, 35, 68, 17, 41, 21, 65, 59, 86, 42, 53, 0, 33, 80, 20},
			//[54,21,73,84,60,18,62,59,89,89,41,55,27,65,94,61,12,76,35,48,0,60,84,9,28,55,4,67,86,33]
		},
	}

	for _, nums := range numss {
		fmt.Println(intersect(nums[0], nums[1]))
	}
}

//如果数字只出现一次的情况
func intersection(nums1 []int, nums2 []int) []int {
	visited := map[int]bool{}
	for _, n := range nums1 {
		visited[n] = true
	}
	var output []int
	for _, n := range nums2 {
		if visited[n] {
			output = append(output, n)
			visited[n] = false
		}
	}
	return output
}

//每个数字尽可能多出现多次数的情况
func intersect(nums1 []int, nums2 []int) []int {
	visited := map[int]int{}
	for _, n := range nums1 {
		val := visited[n]
		val++
		visited[n] = val
	}
	var output []int
	for _, n := range nums2 {
		if val, ok := visited[n]; ok && val != 0 {
			output = append(output, n)
			val--
			visited[n] = val
		}
	}
	return output
}
