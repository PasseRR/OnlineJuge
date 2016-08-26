package algorithms

import (
	"strings"
)

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var result string = ""
	data := "#" + strings.Join(strings.Split(s, ""), "#") + "#"
	length := len(data)
	for index, _ := range data {
		for i := 1; ; i++ {
			// 以每个字符串为旋转中心, 向左右扫描, 直到有对称的字符不同/越界
			if index-i < 0 ||
				i+index+1 > length ||
				data[index-i:index-i+1] != data[i+index:i+index+1] {

				temp := strings.Replace(data[index-i+1:index+i], "#", "", -1)
				if len(temp) > len(result) {
					result = temp
				}
				break
			}
		}
	}

	return result
}
