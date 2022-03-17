package P102

import (
	"runtime/debug"
	"testing"
)

func TestName(t *testing.T) {
	debug.FreeOSMemory()
	levelOrder(nil)
}

type MyNode struct {
	Val  interface{}
	Next *MyNode
	Prev *MyNode
}

type Queue struct {
	len  int
	head *MyNode
	tail *MyNode
}

func NewQueue() *Queue {
	q := new(Queue)
	q.len = 0
	head := new(MyNode)
	tail := new(MyNode)
	head.Next = tail
	tail.Prev = head

	q.head = head
	q.tail = tail
	return q
}

func (q *Queue) add(val interface{}) {
	node := &MyNode{
		Val:  val,
		Next: q.tail,
		Prev: q.tail.Prev,
	}
	node.Prev.Next = node
	node.Next.Prev = node
	q.len++
}

func (q *Queue) poll() interface{} {
	if q.size() <= 0 {
		return nil
	}
	node := q.head.Next
	q.head.Next = q.head.Next.Next
	q.head.Next.Prev = q.head
	q.len--
	return node.Val
}
func (q *Queue) peek() interface{} {
	if q.size() <= 0 {
		return nil
	}
	return q.head.Next.Val
}
func (q *Queue) size() int {
	return q.len
}

func levelOrder(root *TreeNode) [][]int {
	queue := NewQueue()
	var ans [][]int
	if root != nil {
		queue.add(root)
	}
	for queue.size() != 0 {
		n := queue.size()
		var list []int
		for i := 0; i < n; i++ {
			node := queue.poll().(*TreeNode)
			list = append(list, node.Val)
			if node.Left != nil {
				queue.add(node.Left)
			}
			if node.Right != nil {
				queue.add(node.Right)
			}
		}
		ans = append(ans, list)

	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
