package main
import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	len := len(nums)
	for i := 0; i < len; i ++ {
		for j := i + 1; j < len; j ++ {
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("%v", twoSum(nums, target))
}