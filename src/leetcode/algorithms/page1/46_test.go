package page1

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	fmt.Printf("%v\n", permute([]int{1, 2, 3}))
	fmt.Printf("%v\n", permute([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", permute([]int{5, 4, 6, 2}))
	fmt.Printf("%v\n", permute([]int{1, 2, 3, 4, 5}))
}
