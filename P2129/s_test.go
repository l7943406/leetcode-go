package P2129

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(capitalizeTitle("First leTTeR of EACH Word"))
}
func capitalizeTitle(title string) string {
	titleRune := []rune(title)
	for i := 0; i < len(titleRune); i++ {
		// 如果是首字母
		if i-1 < 0 || titleRune[i-1] == ' ' {
			// 如果单词长度大于2
			if i+2 < len(titleRune) && titleRune[i+1] != ' ' && titleRune[i+2] != ' ' {
				// 如果首字母是小写
				if titleRune[i] <= 'z' && titleRune[i] >= 'a' {
					titleRune[i] = titleRune[i] - ('a' - 'A')
				}
				continue
			}
		}
		// 如果非首字母，是大写，改成小写
		if titleRune[i] <= 'Z' && titleRune[i] >= 'A' {
			titleRune[i] = titleRune[i] + ('a' - 'A')
		}
	}
	return string(titleRune)
}
