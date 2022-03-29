package P20

import (
	"testing"
)

func Test(t *testing.T) {
	isValid("()[]{}")
}

func isValid(s string) bool {
	var stack []byte
	top := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			if top+1 >= len(stack) {
				stack = append(stack, s[i])
			} else {
				stack[top+1] = s[i]
			}
			top++
		} else {
			if top >= 0 && (stack[top]+1 == s[i] || stack[top]+2 == s[i]) {
				top--
			} else {
				return false
			}
		}
	}
	return top == -1
}
