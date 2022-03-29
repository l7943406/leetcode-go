package P160

import "testing"

func Test(t *testing.T) {
	getIntersectionNode(nil, nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := headA, headB
	for la != nil && lb != nil {
		if la == lb {
			return la
		}
		la = la.Next
		lb = lb.Next
		if la == nil && lb == nil {
			return nil
		}
		if la == nil {
			la = headB
		}
		if lb == nil {
			lb = headA
		}
	}
	return nil
}
