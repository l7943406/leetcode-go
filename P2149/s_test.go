package P2149

import "testing"

func Test(t *testing.T) {
	rearrangeArray(nil)
}

func rearrangeArray(nums []int) []int {
	target := make([]int, len(nums))

	l, r := 0, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			target[l] = nums[i]
			l += 2
		} else {
			target[r] = nums[i]
			r += 2
		}
	}

	return target
}
