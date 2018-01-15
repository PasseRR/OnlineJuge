package fifty

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	min := math.MaxInt32
	var shortest string
	for _, value := range strs {
		if len(value) < min {
			min = len(value)
			shortest = value
		}
	}

	if len(shortest) == 0 {
		return shortest
	}

	for index, _ := range shortest {
		for _, str := range strs {
			if shortest[0:index+1] != string(str)[0:index+1] {
				return shortest[0:index]
			}
		}
	}

	return shortest
}
