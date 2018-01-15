package fifty

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 9}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 9}, 7))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 9}, 9))
}
