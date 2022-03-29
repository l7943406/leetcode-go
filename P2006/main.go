package main

func main() {

}

func countKDifference(nums []int, k int) (ans int) {
	m := make(map[int]int)
	for _, num := range nums {
		ans += m[num-k] + m[num+k]
		m[num]++
	}
	return ans
}
