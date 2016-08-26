package algorithms

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Printf("%v\n", divide(10, 0))
	fmt.Printf("%v\n", divide(10, 3))
	fmt.Printf("%v\n", divide(10, 2))
	fmt.Printf("%v\n", divide(-1, 1))
	fmt.Printf("%v\n", divide(1, -1))
	fmt.Printf("%v\n", divide(-1, -1))
	fmt.Printf("%v\n", divide(2147483647, 1))
	fmt.Printf("%v\n", math.MaxInt32)
}
