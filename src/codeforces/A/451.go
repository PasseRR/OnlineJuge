package main

import (
	"fmt"
)

// 每选择一个交叉点，去点交叉点的两根木棍
// 即可用的木棍是n，m中较小的一个
func main() {
	var (
		n, m, num int
	)
	fmt.Scanf("%d %d\n", &n, &m)
	if n <= m {
		num = n
	} else {
		num = m
	}
	if num%2 == 0 {
		fmt.Println("Malvika")
	} else {
		fmt.Printf("Akshat")
	}
}
