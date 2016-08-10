package main
import (
	"fmt"
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
		result += x%10
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

func main() {
	fmt.Printf("%v\n", math.MaxInt32)
	fmt.Printf("%v\n", reverse(1534236469))
}
