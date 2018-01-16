package main

import "fmt"

func main() {
	var (
		k    int
		n    int64
		w    int
		cost int64
	)
	fmt.Scanf("%d %d %d", &k, &n, &w)
	for i := 1; i <= w; i++ {
		cost += int64(i * k)
	}
	if cost <= n {
		fmt.Println("0")
	} else {
		fmt.Println(cost - n)
	}
}
