package main

func main() {

}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	i := start
	j := end
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
