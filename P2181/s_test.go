package P2181

import "testing"

func Test(t *testing.T) {
	mergeNodes(nil)
	mergeNodes1(nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 空间复杂度O（n）解法
func mergeNodes(head *ListNode) *ListNode {
	ans := new(ListNode)
	last := ans

	t := head
	num := 0

	for ; t != nil; t = t.Next {
		if t.Val == 0 {
			if num != 0 {
				node := new(ListNode)
				node.Val = num
				last.Next = node
				last = node
				num = 0
			}
			continue
		}
		num += t.Val
	}

	return ans.Next
}

// 空间复杂度O（1）解法
func mergeNodes1(head *ListNode) *ListNode {
	ans := new(ListNode)
	last := ans

	t := head
	num := 0

	for ; t != nil; t = t.Next {
		if t.Val == 0 {
			if num != 0 {
				t.Val = num
				last.Next = t
				last = t
				num = 0
			}
			continue
		}
		num += t.Val
	}
	last.Next = nil
	return ans.Next
}
