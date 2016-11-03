package algorithms

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
