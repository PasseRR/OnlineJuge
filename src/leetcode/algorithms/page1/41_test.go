package page1

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	fmt.Printf("%v\n", firstMissingPositive([]int{10, 1, 2, 7, 6, 1, 5}))
}
