package main

import "fmt"

func main() {
	var (
		n      int
		height = 0
	)
	fmt.Scanln(&n)
	for {
		sum := sum(height + 1)
		if n-sum >= 0 {
			n -= sum
			height++
		} else {
			break
		}
	}

	fmt.Println(height)
}

func sum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}
