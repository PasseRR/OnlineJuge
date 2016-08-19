package leetcode

import (
	"math"
)

func reverse(x int) int {
	flg := false
	if x < 0 {
		flg = true
		x = -x
	}
	var result int
	for x > 0 {
		result *= 10
		result += x % 10
		x /= 10
	}

	if flg {
		result = -result
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result
}
