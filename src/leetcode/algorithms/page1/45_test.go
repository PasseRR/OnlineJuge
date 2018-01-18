package page1

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1, 2, 3}))
	fmt.Println(jump([]int{2, 0, 2, 0, 1}))
	fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
}
