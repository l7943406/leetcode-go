package P206

import "testing"

func Test(t *testing.T) {
	reverseList(nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		curNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = curNext
	}
	return prev
}
