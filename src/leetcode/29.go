package leetcode

import "math"

func divide(dividend int, divisor int) int {
	// 正负标识 true正数 false负数
	flg := true
	if divisor < 0 {
		divisor = -divisor
		flg = !flg
	}
	if dividend < 0 {
		dividend = -dividend
		flg = !flg
	}

	result := 0
	// 判断一个数字是否是2的n次方
	var isPowerOf2 = func(num int) bool {
		return num&(num-1) == 0
	}
	if isPowerOf2(divisor) {
		// 若是2的n次方 直接右移 通过对数函数获得指数
		result = dividend >> uint(math.Log2(float64(divisor)))
	} else {
		// 通过减法来算出商
		for dividend >= 0 {
			dividend -= divisor
			if dividend >= 0 {
				result++
			}
		}
	}

	if !flg {
		result = -result
	}

	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	if result < math.MinInt32 {
		result = math.MinInt32
	}

	return result
}
