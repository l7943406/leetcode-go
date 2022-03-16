package P215

import (
	"container/heap"
	"math/rand"
	"runtime/debug"
	"testing"
	"time"
)

func Test(t *testing.T) {
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	println(findKthLargestS([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() interface{} {
	arr := *h
	return arr[len(arr)-1]
}

func (h *IntHeap) Tail() interface{} {
	arr := *h
	return arr[0]
}

// 堆排 o(nLogK)
func findKthLargest(nums []int, k int) int {
	debug.SetGCPercent(-1)
	h := new(IntHeap)
	heap.Init(h)
	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v > h.Tail().(int) {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	return h.Tail().(int)
}

// 快排剪枝o(n)
func findKthLargestS(nums []int, k int) int {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	sort(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k]
}

// 快排
func sort(nums []int, start int, end int, index int) {
	if end <= start {
		return
	}
	l := start
	r := end
	targetIndex := rand.Intn(end-start+1) + start
	nums[start], nums[targetIndex] = nums[targetIndex], nums[start]
	target := nums[start]

	for l < r {
		for l < r && nums[r] >= target {
			r--
		}
		for l < r && nums[l] <= target {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[start], nums[r] = nums[r], nums[start]

	if r > index {
		sort(nums, start, r-1, index)
	}
	if r < index {
		sort(nums, r+1, end, index)
	}

}
