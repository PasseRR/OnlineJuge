package main

import "fmt"

func generateParenthesis(n int) []string {
	const pair string = "()"
	result := []string{}
	var exist = func(value string) bool {
		for _, item := range result {
			if item == value {
				return true
			}
		}

		return false
	}
	if n == 1 {
		result = append(result, pair)
	} else {
		temp := generateParenthesis(n - 1)
		for _, item := range temp {
			for i, len := -1, len(item); i <= len; i++ {
				var solution string
				if i == -1 { // 在字符串开头加上一个pair
					solution = pair + item
				} else if i == len { // 在字符串末尾加上一个pair
					solution = item + pair
				} else { // 每隔一个字符加上一个pair
					solution = item[0:i] + pair + item[i:len]
				}
				// 判断解决方案是否重复
				if !exist(solution) {
					result = append(result, solution)
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v", generateParenthesis(4))
}
