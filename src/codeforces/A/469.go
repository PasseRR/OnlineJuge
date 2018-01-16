package main

import "fmt"

func main() {
	var (
		n      int
		px, py int
		a      int
		level  [100]int
		flg    bool
	)
	fmt.Scanln(&n)
	fmt.Scanf("%d", &px)
	for i := 0; i < px; i++ {
		fmt.Scanf("%d", &a)
		level[a-1] = 1
	}
	fmt.Scanln()
	fmt.Scanf("%d", &py)
	for i := 0; i < py; i++ {
		fmt.Scanf("%d", &a)
		level[a-1] = 1
	}
	for i := 0; i < n; i++ {
		if level[i] != 1 {
			flg = true
			break
		}
	}

	if flg {
		fmt.Println("Oh, my keyboard!")
	} else {
		fmt.Println("I become the guy.")
	}
}
