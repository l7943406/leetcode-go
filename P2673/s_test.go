package P2130

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(minIncrements(7, []int{
		1,
		5, 2,
		2, 3, 3, 1}))

}
func minIncrements(n int, cost []int) int {
	if n == 0 || n == 1 {
		return 0
	}

	for i := 1; i < n; i++ {
		cost[i] += cost[(i-1)/2]
	}

	length := (n + 1) / 2

	maxCost := cost[n-1]
	for i := n - length; i < n; i++ {
		if cost[i] > maxCost {
			maxCost = cost[i]
		}
	}
	result := 0

	for i := n - length; i < n; i++ {
		cost[i] = maxCost - cost[i]
		result += cost[i]
	}

	n = n - length
	length /= 2
	for n > 0 {
		result2 := result
		for i := n - length; i < n; i++ {
			if cost[(i+1)*2-1] > cost[(i+1)*2] {
				cost[i] = cost[(i+1)*2]
				result -= cost[(i+1)*2]
			} else {
				cost[i] = cost[(i+1)*2-1]
				result -= cost[(i+1)*2-1]
			}
		}
		if result2 == result {
			return result2
		}
		n = n - length
		length /= 2
	}

	return result
}
