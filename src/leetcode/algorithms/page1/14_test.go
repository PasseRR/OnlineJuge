package page1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Logf("%v", longestCommonPrefix([]string{"aab", "aac"}))
}
