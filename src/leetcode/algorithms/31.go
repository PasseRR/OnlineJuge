package algorithms

import (
	"math"
)

func nextPermutation(nums []int) {
	// 是否是降序排列
	flg := true
	len := len(nums)

	// 从尾部开始往前遍历 与前一个元素比较 找到第一个小于当前元素的位置
	for i := len - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			flg = false
			min := math.MaxInt32
			index := -1

			// 找到大于nums[i]最小的元素 并与当前位置交换位置
			for j := i + 1; j < len; j++ {
				if nums[j] > nums[i] && nums[j] < min {
					index = j
					min = nums[j]
				}
			}
			nums[i], nums[index] = nums[index], nums[i]

			// 将i+1到len-1升序排序
			for k := i + 1; k < len; k++ {
				for m := k + 1; m < len; m++ {
					if nums[k] > nums[m] {
						nums[k], nums[m] = nums[m], nums[k]
					}
				}
			}
			break
		}
	}

	if flg {
		// 若是降序排列的 返回一个升序排列
		for i := 0; i < len/2; i++ {
			nums[i], nums[len-1-i] = nums[len-1-i], nums[i]
		}
	}
}
