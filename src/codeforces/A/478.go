package main

import "fmt"

func main() {
	var (
		c1, c2, c3, c4, c5 int
		sum                int
	)
	fmt.Scanln(&c1, &c2, &c3, &c4, &c5)
	sum = c1 + c2 + c3 + c4 + c5
	if sum%5 != 0 || sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum / 5)
	}
}
