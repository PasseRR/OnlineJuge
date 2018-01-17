package page1

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	fmt.Printf("%v\n", trap([]int{0, 2, 1, 0, 1, 3, 2, 1, 2}))
	fmt.Printf("%v\n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
