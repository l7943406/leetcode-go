package main

func countGoodRectangles(rectangles [][]int) int {
	m := map[int]int{}
	max := 0
	for _, rectangle := range rectangles {
		r := min(rectangle[0], rectangle[1])
		m[r]++
		if r > max {
			max = r
		}
	}
	return m[max]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
