package page1

import (
	"math"
)

func trap(height []int) int {
	// 取左右两边第一个挡板向中间移动
	// 在没有遇到高于两个挡板值之前装水量为矮挡板减去当前挡板高度
	sum, start, end := 0, 0, len(height)-1
	for start < end {
		min := int(math.Min(float64(height[start]), float64(height[end])))
		if height[start] == min {
			for start < end && height[start] <= min {
				sum += min - height[start]
				start++
			}
		} else {
			for end > start && height[end] <= min {
				sum += min - height[end]
				end--
			}
		}
	}

	return sum
}
