package algorithms

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}

	fmt.Printf("%v", removeDuplicates(nums))
}
