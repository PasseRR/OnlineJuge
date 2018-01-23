package page1

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	fmt.Printf("%v\n", permuteUnique([]int{1, 1, 2}))
	fmt.Printf("%v\n", permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}
