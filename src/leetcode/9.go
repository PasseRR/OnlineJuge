package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// 将数字转为字符串进行比较是否是回文字符串
	var isPanlindromeString = func(str string) bool {
		length := len(str)
		if length > 1 {
			for i := 0; i <= length/2; i++ {
				if str[i:i+1] != str[length-i-1:length-i] {
					return false
				}
			}
		}

		return true
	}

	return isPanlindromeString(strconv.Itoa(x))
}

func main() {
	fmt.Printf("%v\n", isPalindrome(12321))
}
