package P1588

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum * (len(arr) + 1) / 2
}
