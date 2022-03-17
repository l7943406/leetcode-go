package P25

import "testing"

func Test(t *testing.T) {
	reverseKGroup(nil, 0)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//k个一组反转链表
//原始链表
//ans->0->1->2->3->4->5->nil
// ^   ^     ^  ^
// l1  l2    l3 l4

//反转[l2 , l3]之间的区域，转换为
//ans->2->1->0->3->4->5->nil
// ^   ^     ^  ^
// l1  l3    l2 l4

func reverseKGroup(head *ListNode, k int) *ListNode {
	ans := new(ListNode)
	ans.Next = head

	t := ans
	for {
		l1 := t
		l2 := t.Next
		l3 := t
		for i := 0; i < k; i++ {
			l3 = l3.Next
			if l3 == nil {
				return ans.Next
			}
		}
		l4 := l3.Next

		l3.Next = nil
		reverseList(l2)
		l1.Next = l3
		l2.Next = l4
		t = l2
	}
}

//反转单链表
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
