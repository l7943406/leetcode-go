package P88

import "testing"

func Test(t *testing.T) {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	top := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[top] = nums1[m]
			m--
		} else {
			nums1[top] = nums2[n]
			n--
		}
		top--
	}
	for n >= 0 {
		nums1[n] = nums2[n]
		n--
	}

}
