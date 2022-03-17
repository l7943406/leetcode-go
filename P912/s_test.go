package P912

import (
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	sortArray(nil)
}
func sortArray(nums []int) []int {
	//随机种子
	rand.Seed(time.Now().UnixNano())
	sort(nums, 0, len(nums)-1)
	return nums
}
func sort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	randomIndex := rand.Intn(end-start+1) + start
	nums[start], nums[randomIndex] = nums[randomIndex], nums[start]
	targetNum := nums[start]
	l := start
	r := end
	for l < r {
		for l < r && nums[r] >= targetNum {
			r--
		}
		for l < r && nums[l] <= targetNum {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[r], nums[start] = nums[start], nums[r]
	sort(nums, start, r-1)
	sort(nums, r+1, end)

}
