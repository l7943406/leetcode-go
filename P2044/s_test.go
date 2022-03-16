package P2044

import (
	"testing"
)

func Test(t *testing.T) {
	countMaxOrSubsets([]int{3, 2, 1, 5})
}

func countMaxOrSubsets(nums []int) int {
	max := 0
	for _, v := range nums {
		max |= v
	}

	return 0
}
