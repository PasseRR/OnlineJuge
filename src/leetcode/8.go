package main
import (
	"strconv"
	"fmt"
	"strings"
	"math"
)

func myAtoi(str string) int {
	// flg 截取字符串标识 当前字符串是否包含不为空的字符串
	// start 不为空字符串开始索引
	flg, start := false, 0
	digits := "0123456789"
	for index, value := range str {
		if string(value) == " " {
			if flg {
				// 若遇到非空字符后再遇到空字符 直接跳出并截取字符串
				str = str[start : index]
				break
			}else {
				continue
			}
		}else {
			if !flg {
				start = index
				flg = true
			}

			// 第一个非空字符是 + -
			if index == start &&
			(string(value) == "-" || string(value) == "+") {
				continue
			}

			// 当遇到非数字字符
			if !strings.ContainsAny(string(value), digits) {
				str = str[start : index]
				break
			}

			// 遍历到字符串最后
			if index == len(str)-1 {
				str = str[start : len(str)]
			}
		}
	}

	i, _ := strconv.Atoi(str)
	if i > math.MaxInt32 {
		i = math.MaxInt32
	}

	if i < math.MinInt32 {
		i = math.MinInt32
	}
	return i
}

func main() {
	fmt.Printf("%v\n", myAtoi("123  456"))
}
