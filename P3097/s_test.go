package P3095

import (
	"math"
	"testing"
)

/*
给你一个 非负 整数数组 nums 和一个整数 k 。

如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。

请你返回 nums 中 最短特别非空
子数组
的长度，如果特别子数组不存在，那么返回 -1 。

示例 1：

输入：nums = [1,2,3], k = 2

输出：1

解释：

子数组 [3] 的按位 OR 值为 3 ，所以我们返回 1 。

示例 2：

输入：nums = [2,1,8], k = 10

输出：3

解释：

子数组 [2,1,8] 的按位 OR 值为 11 ，所以我们返回 3 。

示例 3：

输入：nums = [1,2], k = 0

输出：1

解释：

子数组 [1] 的按位 OR 值为 1 ，所以我们返回 1 。

提示：

1 <= nums.length <= 2 * 105
0 <= nums[i] <= 109
0 <= k <= 109
*/
func TestName(t *testing.T) {

}

func minimumSubarrayLength(nums []int, k int) int {
	bitsMap := make([]int, 32)

	result := -1
	i := 0
	j := 0

	for j < len(nums) {
		if result == 1 {
			return 1
		}
		or(bitsMap, nums[j])
		j++
		for check(bitsMap, k) {
			if result == -1 || j-i < result {
				result = j - i
			}
			if i < j-1 {
				subOr(bitsMap, nums[i])
				i++
			} else {
				break
			}
		}
	}

	return result
}

func or(bitsMap []int, num int) {
	for l := 0; l < 30; l++ {
		if num&(1<<l) != 0 {
			bitsMap[l]++
		}
	}
}

func subOr(bitsMap []int, num int) {
	for l := 0; l < 30; l++ {
		if num&(1<<l) != 0 {
			bitsMap[l]--
		}
	}
}

func check(bitsMap []int, k int) bool {
	num := 0
	for i := 0; i < len(bitsMap); i++ {
		if bitsMap[i] == 0 {
			continue
		}
		num = num + 1<<i
	}
	return num >= k
}

func minimumSubarrayLength2(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr >= k {
			// ans = min(ans, right-left+1)
			left++
			if bottom < left {
				// 重新构建一个栈
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
