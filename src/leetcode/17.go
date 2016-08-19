package leetcode

func letterCombinations(digits string) []string {
	digitMapLetter := make(map[string]string)
	digitMapLetter["2"] = "abc"
	digitMapLetter["3"] = "def"
	digitMapLetter["4"] = "ghi"
	digitMapLetter["5"] = "jkl"
	digitMapLetter["6"] = "mno"
	digitMapLetter["7"] = "pqrs"
	digitMapLetter["8"] = "tuv"
	digitMapLetter["9"] = "wxyz"
	result := []string{}
	combination := func(data []string, str string) []string {
		var comb []string
		// 第一次组合 分别把数字代表的几个字母放入数组
		if len(data) == 0 {
			for _, value := range str {
				comb = append(comb, string(value))
			}
		} else {
			// 分别和当前已经组合的数组再次组合 当前数组和数字代表的字母依次组合
			for _, value := range str {
				for _, item := range data {
					comb = append(comb, item+string(value))
				}
			}
		}

		return comb
	}

	for _, value := range digits {
		result = combination(result, digitMapLetter[string(value)])
	}

	return result
}
