package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{54, 21, 73, 84, 60, 18, 62, 59, 89, 89, 41, 55, 27, 65, 94, 61, 12, 76, 35, 48, 0, 60, 84, 9, 28, 55, 4, 67, 86, 33}
	b := []int{18, 73, 59, 89, 84, 84, 48, 84, 54, 62, 67, 27, 60, 0, 61, 94, 84, 55, 55, 60, 76, 35, 84, 28, 4, 9, 86, 12, 89, 41, 21, 65, 33}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Println(a)
	fmt.Println(b)
}
