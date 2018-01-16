package page1

import "sort"

// nums 长度可能为0
// 存在重复的数字
func firstMissingPositive(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	// 数组长度为0或者最小数大于1或者最大数小于等于0
	if length == 0 || nums[0] > 1 || nums[length-1] <= 0 {
		return 1
	}

	var positiveIndex int
	// 找到第一个大于零数字的索引
	for i, value := range nums {
		if value > 0 {
			positiveIndex = i
			break
		}
	}

	// 若非零开始索引数字不为1 则返回1
	if nums[positiveIndex] > 1 {
		return 1
	}

	for i := positiveIndex + 1; i < length; i++ {
		// 相等或者相差1
		if nums[i-1] == nums[i] || nums[i-1]+1 == nums[i] {
			continue
		}

		if nums[i-1]+1 != nums[i] {
			return nums[i-1] + 1
		}
	}

	return nums[length-1] + 1
}
