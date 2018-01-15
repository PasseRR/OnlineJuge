package fifty

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	}
	maxLen := 1
	var subString string
	for i := 0; i < length; i++ {
		subString = ""
		for j := i; j < length; j++ {
			if !strings.Contains(subString, s[j:j+1]) {
				subString = s[i : j+1]
			} else {
				currentLen := len(subString)
				if currentLen > maxLen {
					maxLen = currentLen
				}
				break
			}
		}
		currentLen := len(subString)
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
