package P1971

func findCenter(edges [][]int) int {
	i, j := edges[0][0], edges[0][1]
	m, n := edges[1][0], edges[1][1]
	if i == m || i == n {
		return i
	} else {
		return j
	}
}
