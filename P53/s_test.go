package P53

import "testing"

func Test(t *testing.T) {
	maxSubArray(nil)
}

func maxSubArray(nums []int) int {
	ans := nums[0]
	num := 0
	for _, v := range nums {
		num += v
		ans = max(ans, num)
		if num < 0 {
			num = 0
		}
	}

	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
