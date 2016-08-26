package algorithms

func romanToInt(s string) int {
	romans := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I",
	}
	nums := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}
	result := 0
	// len1 字符串长度
	j, len1 := 0, len(s)
	// len2 对照数组长度
	for i, len2 := 0, len(romans); i < len2; i++ {
		// 罗马数字长度
		len3 := len(romans[i])
		for j < len1 && j+len3 <= len1 && s[j:j+len3] == romans[i] {
			result += nums[i]
			j += len3
		}
	}

	return result
}
