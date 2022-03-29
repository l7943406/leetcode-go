package main

func main() {
	imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}})
}
func imageSmoother(img [][]int) [][]int {
	ans := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		ans[i] = make([]int, len(img[i]))
		for j := 0; j < len(img[i]); j++ {
			ans[i][j] = getAvg(&img, i, j)
		}
	}
	return ans

}

func getAvg(img *[][]int, x int, y int) int {
	sum := 0
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < len(*img) && j >= 0 && j < len((*img)[0]) {
				sum += (*img)[i][j]
				count++
			}
		}
	}
	return sum / count
}
