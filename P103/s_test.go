package P103

import "testing"

func Test(t *testing.T) {
	zigzagLevelOrder(&TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	deque := NewDeque()
	deque.AddLast(root)
	lToR := true
	for deque.Size() > 0 {
		n := deque.Size()
		var list []int
		for i := 0; i < n; i++ {
			node := deque.PollFirst()
			if node.Left != nil {
				deque.AddLast(node.Left)
			}
			if node.Right != nil {
				deque.AddLast(node.Right)
			}
			if lToR {
				list = append(list, node.Val)
			} else {
				list = append([]int{node.Val}, list...)
			}
		}
		ans = append(ans, list)
		lToR = !lToR
	}

	return ans
}

type Deque struct {
	Len        int
	Head, Tail *MyNode
}

type MyNode struct {
	Val        *TreeNode
	next, prev *MyNode
}

func (q *Deque) AddFirst(val *TreeNode) {
	node := &MyNode{
		Val:  val,
		next: q.Head.next,
		prev: q.Head,
	}
	node.next.prev = node
	node.prev.next = node
	q.Len++
}
func (q *Deque) AddLast(val *TreeNode) {
	node := &MyNode{
		Val:  val,
		next: q.Tail,
		prev: q.Tail.prev,
	}
	node.next.prev = node
	node.prev.next = node
	q.Len++
}
func (q *Deque) PollLast() *TreeNode {
	if q.Len <= 0 {
		return nil
	}
	node := q.Tail.prev
	node.next.prev = node.prev
	node.prev.next = node.next
	q.Len--
	return node.Val
}

func (q *Deque) PollFirst() *TreeNode {
	if q.Len <= 0 {
		return nil
	}
	node := q.Head.next
	node.next.prev = node.prev
	node.prev.next = node.next
	q.Len--
	return node.Val
}
func (q *Deque) Size() int {
	return q.Len
}
func NewDeque() *Deque {
	head, tail := new(MyNode), new(MyNode)
	head.next = tail
	tail.prev = head
	return &Deque{
		Len:  0,
		Head: head,
		Tail: tail,
	}
}
