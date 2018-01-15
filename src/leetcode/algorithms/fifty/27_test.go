package fifty

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{2}

	fmt.Printf("%v", removeElement(nums, 4))
}
