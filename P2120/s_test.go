package P2120

import "testing"

func Test(t *testing.T) {
	executeInstructions(3, []int{0, 1}, "RRDDLU")
}

func executeInstructions(n int, startPos []int, s string) []int {
	m := map[uint8][]int{
		'L': {0, -1},
		'R': {0, 1},
		'D': {1, 0},
		'U': {-1, 0},
	}
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ans[i] = solve(m, n, startPos[0], startPos[1], s, i)
	}
	return ans
}

func solve(m map[uint8][]int, n int, x int, y int, s string, startIndex int) int {

	count := 0
	for i := startIndex; i < len(s); i++ {
		x += m[s[i]][0]
		y += m[s[i]][1]
		if x < 0 || x >= n || y < 0 || y >= n {
			break
		}
		count++
	}
	return count
}
