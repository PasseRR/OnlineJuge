package main
import (
	"fmt"
	"sort"
)

func main() {
	var (
		n int
		num int
		nums = make([]int, 0, 100)
	)
	fmt.Scanln(&n)
	for i := 0; i < n; i ++ {
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}
	sort.Ints(nums)
	for i := 0; i < n-1; i ++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Print(nums[n-1])
}
