# 题解

## 一、删除排序数组中的重复项

给定一个排序数组，你需要在 `原地` 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 `原地` 修改输入数组 并在使用 `O(1)` 额外空间的条件下完成。

> 基本思路

因为是已经排序好的，所以之前出现的过数字只会连续出现，不会重复在后面重复出现。 

可以使用快慢指针的方式，一个指针遍历数组，另一个指针记录未重复出现数字的在原数组的位置。

快指针遍历完数组则结束

慢指针在遇到与当前值不一样的数字时，将该数字记录到当前位置，并向前移动一步。

```go
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
```