package P2130

import (
	"testing"
)

func Test(t *testing.T) {
	head := new(ListNode)
	head.Next = new(ListNode)
	head.Next.Next = new(ListNode)
	head.Next.Next.Next = new(ListNode)
	head.Val = 5
	head.Next.Val = 4
	head.Next.Next.Val = 1
	head.Next.Next.Next.Val = 2
	pairSum(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = head
	fast = reserve(fast)
	maxNum := 0
	for slow != nil {
		maxNum = max(maxNum, slow.Val+fast.Val)
		slow = slow.Next
		fast = fast.Next
	}

	return maxNum
}
func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func reserve(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		curNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = curNext
	}
	return prev
}
