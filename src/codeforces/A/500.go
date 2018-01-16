package main

import "fmt"

func main() {
	var (
		n, t  int
		index = 1
	)
	fmt.Scanln(&n, &t)
	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i, ai := range a {
		if (i + 1) == index {
			index += ai
			if index == t {
				fmt.Println("Yes")
				break
			} else if index > t {
				fmt.Println("No")
				break
			}
		}
	}
}
