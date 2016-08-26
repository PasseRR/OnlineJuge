package algorithms

import (
	"bytes"
)

func convert(s string, numRows int) string {
	length := len(s)
	if length == 0 || length == 1 || numRows == 1 {
		return s
	}
	// golang 字符串拼接
	var buffer bytes.Buffer
	// z字直线部分索引高度 从0开始
	height := numRows - 1
	// 分析z字形字符串与返回字符串位置关系
	for i := 0; i < numRows; i++ {
		for j := i; j < length; j += height * 2 {
			// 第一行和最后一行只有z字直线部分
			buffer.WriteString(s[j : j+1])
			// 斜线每行位移大小
			offset := j + 2*(height-i)
			// 中间行都有z字斜线部分
			if i%height != 0 && offset < length {
				buffer.WriteString(s[offset : offset+1])
			}
		}
	}
	return buffer.String()
}
