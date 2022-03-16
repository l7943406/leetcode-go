package P146

import (
	"runtime/debug"
	"testing"
)

func Test(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)

}

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	head     *Node
	tail     *Node
	m        map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	debug.SetGCPercent(-1)
	head := Node{
		Key:   0,
		Value: 0,
		Next:  nil,
		Prev:  nil,
	}
	tail := Node{
		Key:   0,
		Value: 0,
		Next:  nil,
		Prev:  nil,
	}
	head.Next = &tail
	tail.Prev = &head

	return LRUCache{
		head:     &head,
		tail:     &tail,
		m:        map[int]*Node{},
		capacity: capacity,
	}
}

// 删除最旧元素
func (lru *LRUCache) deleteOldestNode() *Node {
	if len(lru.m) == 0 {
		return nil
	}
	old := lru.tail.Prev
	lru.deleteNode(old)
	return old
}

// 从LRU中删除key
func (lru *LRUCache) deleteNode(node *Node) {
	delete(lru.m, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 添加新元素
func (lru *LRUCache) addNewest(node *Node) {
	// 如果满了就释放空间
	for len(lru.m) >= lru.capacity {
		lru.deleteOldestNode()
	}
	node.Next = lru.head.Next
	node.Prev = lru.head.Next.Prev

	node.Prev.Next = node
	node.Next.Prev = node
	lru.m[node.Key] = node

}

func (lru *LRUCache) Get(key int) int {
	value, ok := lru.m[key]
	if !ok {
		return -1
	}
	lru.deleteNode(value)
	lru.addNewest(value)
	return value.Value
}

func (lru *LRUCache) Put(key int, value int) {
	v, ok := lru.m[key]
	if ok {
		lru.deleteNode(v)
	}
	lru.addNewest(&Node{
		Key:   key,
		Value: value,
		Next:  nil,
		Prev:  nil,
	})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
