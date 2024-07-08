package P1673

import (
	"fmt"
	"testing"
)

// 1673. 找出最具竞争力的子序列
// 提示
// 给你一个整数数组 nums 和一个正整数 k ，返回长度为 k 且最具 竞争力 的 nums 子序列。
//
// 数组的子序列是从数组中删除一些元素（可能不删除元素）得到的序列。
//
// 在子序列 a 和子序列 b 第一个不相同的位置上，如果 a 中的数字小于 b 中对应的数字，那么我们称子序列 a 比子序列 b（相同长度下）更具 竞争力 。 例如，[1,3,4] 比 [1,3,5] 更具竞争力，在第一个不相同的位置，也就是最后一个位置上， 4 小于 5 。
//
//
//
// 示例 1：
//
// 输入：nums = [3,5,2,6], k = 2
// 输出：[2,6]
// 解释：在所有可能的子序列集合 {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]} 中，[2,6] 最具竞争力。
// 示例 2：
//
// 输入：nums = [2,4,3,3,2,4,9,6], k = 4
// 输出：[2,3,3,4]

type NumsInfo struct {
	index, value int
}

func mostCompetitive(nums []int, k int) []int {
	result := make([]int, k)
	index := 0
	for i := 0; i < len(nums); i++ {
		for index > 0 && result[index-1] > nums[i] && len(nums)-i > k-index && index < k {
			index--
		}
		result[index] = nums[i]
		index++
	}

	return result
}

func TestC(t *testing.T) {
	fmt.Println(mostCompetitive([]int{3, 5, 2, 6}, 2))
}
