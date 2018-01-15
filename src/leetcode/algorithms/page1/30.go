package page1

// 考虑words中会有重复单词
func findSubstring(s string, words []string) []int {
	len1 := len(s)
	len2 := len(words)
	// 单词长度 每个单词长度一样
	var lenOfWord int
	if len(words) >= 1 {
		lenOfWord = len(words[0])
	}
	// 索引结果集
	result := []int{}
	// 如果字符串长度小于单词组加起来的长度
	if len1 < len2*lenOfWord {
		return result
	}
	// 单词是否出现在当前遍历子字符串中
	existMap := make(map[string]int)

	// 重置map
	var resetMap = func() {
		for key, _ := range existMap {
			existMap[key] = 0
		}
		for _, value := range words {
			count, exist := existMap[value]
			if exist {
				existMap[value] = count + 1
			} else {
				existMap[value] = 1
			}
		}
	}
	// 所有单词已经在当前遍历中是否已经全部出现
	var isAllExist = func() bool {
		for _, value := range existMap {
			if value != 0 {
				return false
			}
		}

		return true
	}

	resetMap()
	for i := 0; i <= len1-lenOfWord; {
		for j := i; j <= len1-lenOfWord; j += lenOfWord {
			word := s[j : j+lenOfWord]
			// 判断map中是否有该单词
			count, exist := existMap[word]
			if exist {
				if count == 0 {
					break
				} else {
					existMap[word] = count - 1
				}

				if isAllExist() {
					result = append(result, i)
					break
				}
			} else {
				break
			}
		}
		resetMap()
		i++
	}

	return result
}
