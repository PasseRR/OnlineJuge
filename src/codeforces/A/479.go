package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scanf("%d\n%d\n%d", &a, &b, &c)

	if a == 1 && c == 1 {
		fmt.Println(a + b + c)
	} else if a == 1 {
		fmt.Println((a + b) * c)
	} else if c == 1 {
		fmt.Println(a * (b + c))
	} else if b == 1 {
		if a > c {
			fmt.Println(a * (b + c))
		} else {
			fmt.Println((a + b) * c)
		}
	} else {
		fmt.Println(a * b * c)
	}
}
