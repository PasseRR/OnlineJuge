package algorithms

import (
	"fmt"
	"testing"
)

func TestNormalSearch(t *testing.T) {
	fmt.Println(normalSearch([]int{1, 3, 2}, 1))
	fmt.Println(normalSearch([]int{1, 3, 2}, 2))
	fmt.Println(normalSearch([]int{1, 3, 2}, 3))
	fmt.Println(normalSearch([]int{1, 3, 2}, 4))
	fmt.Println(normalSearch([]int{1, 3, 2}, 5))
}
