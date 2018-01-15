package page1

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	len := len(nums)
	var result [][]int
	var exist = func(data []int) bool {
		for _, value := range result {
			if value[0] == data[0] &&
				value[1] == data[1] &&
				value[2] == data[2] {
				return true
			}
		}

		return false
	}
	for i := 0; i < len-2; i++ {
		for j := i + 1; j < len-1; j++ {
			for k := j + 1; k < len; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					item := []int{nums[i], nums[j], nums[k]}

					if !exist(item) {
						result = append(result, item)
					}
				}
			}
		}
	}

	return result
}
