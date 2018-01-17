package page1

import (
	"fmt"
	"math"
)

func multiply(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	nums1, nums2 := make([]int, len1), make([]int, len2)
	result := make([]int, int(math.Max(float64(len1), float64(len2))*2))
	// 将字符串转为数字倒叙放入数组
	for i, j := len1-1, 0; i >= 0; i, j = i-1, j+1 {
		nums1[j] = int(num1[i] - '0')
	}
	for i, j := len2-1, 0; i >= 0; i, j = i-1, j+1 {
		nums2[j] = int(num2[i] - '0')
	}

	// 计算每个位置乘积
	for i, a := range nums1 {
		for j, b := range nums2 {
			// i + j 表示十进制位数 0表示个位一次递增
			result[i+j] += a * b
		}
	}

	// 对计算结果进行进位
	for i, value := range result {
		if value >= 10 {
			result[i+1] += value / 10
			result[i] = value % 10
		}
	}

	// product乘积字符串
	product := ""
	// 倒叙输出result 去掉前置0为结果
	for i := len(result) - 1; i >= 0; i-- {
		if len(product) != 0 || result[i] > 0 {
			product += fmt.Sprintf("%d", result[i])
		}
	}

	if len(product) == 0 {
		product = "0"
	}

	return product
}
