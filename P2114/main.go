package main

func main() {
	mostWordsFound([]string{
		"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much",
	})
}

func mostWordsFound(sentences []string) int {

	ans := 0
	for _, s := range sentences {
		sum := 1
		for _, c := range s {
			if c == ' ' {
				sum++
			}
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
