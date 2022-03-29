package P2194

func cellsInRange(s string) []string {
	cStart, cEnd := s[0], s[3]
	rStart, rEnd := s[1], s[4]
	var ans []string
	for i := cStart; i <= cEnd; i++ {
		for j := rStart; j <= rEnd; j++ {
			ans = append(ans, string([]byte{i, j}))
		}
	}
	return ans
}
