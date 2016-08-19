package leetcode

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	len1, len2 := len(haystack), len(needle)
	if len1 == len2 && len1 == 0 {
		return 0
	}

	if len1 < len2 {
		return -1
	}

	for i := 0; i <= len1-len2; i++ {
		if haystack[i:i+len2] == needle {
			return i
		}
	}

	return -1
}

// use golang lib
func index(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
