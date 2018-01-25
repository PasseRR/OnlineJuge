package page1

const (
	QUESTION_MASK = '?'
	ASTERISK      = '*'
)

func isMatch2(s, p string) bool {
	lengthS, lengthP := len(s), len(p)
	sIndex, pIndex, matchIndex, startIndex := 0, 0, 0, -1
	for sIndex < lengthS {
		if pIndex < lengthP && (p[pIndex] == QUESTION_MASK || p[pIndex] == s[sIndex]) {
			// 匹配字符串索引位置字符为?或与待匹配字符索引位置字符相同
			// 索引递增 继续匹配
			pIndex++
			sIndex++
		} else if pIndex < lengthP && p[pIndex] == ASTERISK {
			// 若当前匹配字符串索引位置字符为*
			// 记录*字符所在索引
			startIndex = pIndex
			// 记录*匹配上的索引
			matchIndex = sIndex
			pIndex++
		} else if startIndex != -1 {
			pIndex = startIndex + 1
			matchIndex++
			sIndex = matchIndex
		} else {
			// 匹配不上
			return false
		}
	}

	// 匹配字符串以*结尾
	for pIndex < lengthP && p[pIndex] == ASTERISK {
		pIndex++
	}

	return pIndex == lengthP
}

// TLE
func isMatch3(s string, p string) bool {
	lengthP, lengthS := len(p), len(s)
	if lengthS == 0 {
		if lengthP == 0 {
			return true
		} else {
			flag := true
			for _, value := range p {
				if value != ASTERISK {
					flag = false
					break
				}
			}

			return flag
		}
	}

	if lengthP == 0 {
		return lengthS == 0
	}

	currentByteP, currentByteS := p[0], s[0]
	if currentByteP == ASTERISK {
		for i := 0; i <= lengthS; i++ {
			if isMatch2(s[i:], p[1:]) {
				return true
			}
		}
	} else if currentByteP == QUESTION_MASK || currentByteP == currentByteS {
		return isMatch2(s[1:], p[1:])
	}

	return false
}
