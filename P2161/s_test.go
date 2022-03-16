package P2161

import "testing"

func Test(t *testing.T) {
	pivotArray(nil, 0)
}

//空间复杂度o（n）
func pivotArray(nums []int, pivot int) []int {
	if len(nums) <= 1 {
		return nums
	}
	target := make([]int, len(nums))

	l, r := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			target[l] = nums[i]
			l++
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > pivot {
			target[r] = nums[i]
			r--
		}
	}
	for i := l; i <= r; i++ {
		target[i] = pivot
	}

	return target
}
