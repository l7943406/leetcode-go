package P121

import (
	"testing"
)

func Test(t *testing.T) {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}

func maxProfit(prices []int) int {
	ans := 0
	min := prices[0]

	for _, v := range prices {
		if v <= min {
			min = v
		} else {
			if v-min > ans {
				ans = v - min
			}
		}
	}
	return ans
}
