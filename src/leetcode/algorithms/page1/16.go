package page1

import (
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	offset := math.MaxInt32
	len := len(nums)
	for i := 0; i < len-2; i++ {
		for j := i + 1; j < len-1; j++ {
			for k := j + 1; k < len; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum == target {
					return target
				} else if sum > target {
					if sum-target < offset {
						result = sum
						offset = sum - target
					}
				} else if sum < target {
					if target-sum < offset {
						result = sum
						offset = target - sum
					}
				}
			}
		}
	}

	return result
}
