package page1

func longestValidParentheses(s string) int {
	// 记录上次匹配失败的地方
	// 默认为没有失败的 则会加上0位置的长度
	lastFail := []int{-1}
	max := 0
	for i, length := 0, len(s); i < length; i++ {
		if s[i:i+1] == "(" {
			lastFail = append(lastFail, i)
		} else {
			length := len(lastFail)
			index := lastFail[length-1]
			// 如果当前栈顶是( 那么当前栈中肯定有2个以上元素
			if index >= 0 && s[index:index+1] == "(" {
				// 弹出栈顶
				lastFail = lastFail[0 : length-1]
				// 取得上一次匹配失败的索引位置
				count := i - lastFail[length-2]
				if count > max {
					max = count
				}
			} else {
				// 如果为空栈或者栈顶为)
				lastFail = append(lastFail, i)
			}
		}
	}

	return max
}
