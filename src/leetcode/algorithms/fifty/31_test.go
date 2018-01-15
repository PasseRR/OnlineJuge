package fifty

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{1, 1, 5})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{1, 3, 2})
	nextPermutation([]int{5, 4, 3, 2, 1})
	nextPermutation([]int{5, 4, 3, 2})
	nextPermutation([]int{1, 2, 3, 4})
}
