package P2125

import "testing"

func Test(t *testing.T) {
	numberOfBeams(nil)
}

func numberOfBeams(bank []string) int {
	ans := 0
	last := 0
	for _, s := range bank {
		num := getCount(s)
		if num != 0 {
			ans += num * last
			last = num
		}
	}

	return ans
}

func getCount(s string) int {
	count := 0
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	return count
}
