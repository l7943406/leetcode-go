package P21

import (
	"testing"
)

func Test(t *testing.T) {
	mergeTwoLists(nil, nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := new(ListNode)
	prev := ans
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			prev = list1
			list1 = list1.Next
			prev.Next = nil
		} else {
			prev.Next = list2
			prev = list2
			list2 = list2.Next
			prev.Next = nil
		}
	}

	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}

	return ans.Next
}
