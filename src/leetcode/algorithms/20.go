package algorithms

func isValid(s string) bool {
	parentheses := []string{}
	for _, value := range s {
		parenthesis := string(value)
		if parenthesis == "(" ||
			parenthesis == "[" ||
			parenthesis == "{" {
			// 开括号入栈
			parentheses = append(parentheses, parenthesis)
		} else {
			len := len(parentheses)
			// 括号是否是开关对应
			if len < 1 ||
				(parenthesis == ")" && parentheses[len-1] != "(") ||
				(parenthesis == "]" && parentheses[len-1] != "[") ||
				(parenthesis == "}" && parentheses[len-1] != "{") {
				return false
			} else {
				// 开括号出栈
				parentheses = parentheses[0 : len-1]
			}
		}
	}

	return len(parentheses) == 0
}
