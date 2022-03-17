package P15

import (
	"runtime/debug"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
	threeSum2([]int{-1, 0, 1, 2, -1, -4})
	threeSum3([]int{-1, 0, 1, 2, -1, -4})
}

//循环 + 二分
func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for l := 0; l < len(nums)-2 && nums[l] <= 0; {
		lNum := nums[l]
		for r := len(nums) - 1; l+1 < r && nums[r] >= 0; {
			rNum := nums[r]
			targetNum := -lNum - rNum
			if binarySearch(nums, targetNum, l+1, r-1) != -1 {
				ans = append(ans, []int{lNum, targetNum, rNum})
			}
			for l+1 < r && nums[r] == rNum {
				r--
			}
		}
		for l < len(nums)-2 && nums[l] == lNum {
			l++
		}
	}
	return ans
}
func binarySearch(nums []int, target int, l int, r int) int {
	for l <= r {
		mid := (l-r)/2 + r
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

//循环 + map
func threeSum2(nums []int) [][]int {
	debug.SetGCPercent(-1)
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	var ans [][]int
	sort.Ints(nums)
	for l := 0; l < len(nums)-2 && nums[l] <= 0; {
		lNum := nums[l]
		for r := len(nums) - 1; l+1 < r && nums[r] >= 0; {
			rNum := nums[r]
			targetNum := -lNum - rNum
			if lNum <= targetNum && targetNum <= rNum {
				count := m[targetNum]
				if lNum == targetNum {
					count--
				}
				if rNum == targetNum {
					count--
				}
				if count > 0 {

					ans = append(ans, []int{lNum, targetNum, rNum})
				}
			}
			for l+1 < r && nums[r] == rNum {
				r--
			}
		}
		for l < len(nums)-2 && nums[l] == lNum {
			l++
		}
	}
	return ans
}

//双指针
func threeSum3(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)

	for l := 0; l < len(nums)-2 && nums[l] <= 0; {
		lNum := nums[l]
		r := len(nums) - 1
		m := l + 1
		for m < r {
			mNum := nums[m]
			for m < r && lNum+mNum+nums[r] > 0 {
				r--
			}
			if m == r {
				break
			}

			if lNum+mNum+nums[r] == 0 {
				ans = append(ans, []int{lNum, mNum, nums[r]})
			}

			for m < r && nums[m] == mNum {
				m++
			}
		}
		for l < len(nums)-2 && nums[l] == lNum {
			l++
		}
	}
	return ans
}
