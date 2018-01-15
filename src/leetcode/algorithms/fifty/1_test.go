package fifty

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("%v\n", twoSum(nums, target))
}
