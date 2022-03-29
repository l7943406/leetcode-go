package P1

import "testing"

func Test(t *testing.T) {
	twoSum(nil, 0)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int)
	for index, value := range nums {
		if _, ok := m[value]; !ok {
			m[value] = []int{}
		}
		m[value] = append(m[value], index)
	}
	for index, value := range nums {
		if len(m[target-value]) > 0 {
			if 2*value == target && len(m[value]) > 1 {
				return []int{m[value][0], m[value][1]}
			}
			if 2*value != target {
				return []int{index, m[target-value][0]}
			}
		}
	}
	return []int{}

}
