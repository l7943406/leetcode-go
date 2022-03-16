package P3

import "testing"

func Test(t *testing.T) {
	lengthOfLongestSubstring("")
}

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	m := map[uint8]interface{}{}
	ans := 0
	for r < len(s) {
		for m[s[r]] != nil {
			delete(m, s[l])
			l++
		}
		m[s[r]] = 0
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
