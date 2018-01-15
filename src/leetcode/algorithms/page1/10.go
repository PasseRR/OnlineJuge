package page1

func isMatch(s string, p string) bool {
	i, j := 0, 0
	len1, len2 := len(s), len(p)
	var offset int
	for i < len1 && j < len2 {
		if j+1 < len2 && p[j+1:j+2] == "*" {
			offset = 2
			for i < len1 && (s[i:i+1] == p[j:j+1] || p[j:j+1] == ".") {
				// aaa a*a匹配情形
				if isMatch(s[i:len1], p[j+offset:len2]) {
					return true
				}
				i++
			}
		} else {
			offset = 1
			if s[i:i+1] == p[j:j+1] || p[j:j+1] == "." {
				i++
			} else {
				return false
			}
		}

		j += offset
	}

	// 如果没有匹配完字符串
	if i < len1 {
		return false
	}

	// 剩余被匹配字符串中是否包含*
	if j < len2 {
		for j < len2 {
			if j+1 < len2 && p[j+1:j+2] == "*" {
				j += 2
			} else {
				return false
			}
		}
	}

	return true
}
