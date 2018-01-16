package main

import "fmt"

func main() {
	var (
		n, m, a, b int
		per        float32
	)
	fmt.Scanln(&n, &m, &a, &b)
	per = float32(b / m)
	if per >= float32(a) {
		fmt.Println(n * a)
	} else {
		if n%m == 0 {
			fmt.Println(n / m * b)
		} else {
			remain := n % m
			if remain*a >= b {
				fmt.Println((n/m + 1) * b)
			} else {
				fmt.Println(remain*a + (n/m)*b)
			}
		}
	}
}
