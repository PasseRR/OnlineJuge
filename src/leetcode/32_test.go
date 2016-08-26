package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("()(()"))
}
