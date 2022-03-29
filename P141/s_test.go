package P141

import "testing"

func Test(t *testing.T) {
	hasCycle(nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
