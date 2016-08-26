package algorithms

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	len := len(nums)
	var result [][]int
	var exist = func(data []int) bool {
		for _, value := range result {
			if value[0] == data[0] &&
				value[1] == data[1] &&
				value[2] == data[2] &&
				value[3] == data[3] {
				return true
			}
		}

		return false
	}
	for i := 0; i < len-3; i++ {
		for j := i + 1; j < len-2; j++ {
			for k := j + 1; k < len-1; k++ {
				for m := k + 1; m < len; m++ {
					if nums[i]+nums[j]+nums[k]+nums[m] == target {
						item := []int{nums[i], nums[j], nums[k], nums[m]}

						if !exist(item) {
							result = append(result, item)
						}
					}
				}
			}
		}
	}

	return result
}
