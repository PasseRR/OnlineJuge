package main

import "fmt"

func main() {
	var (
		n, m  int
		index int
	)
	fmt.Scanln(&n, &m)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			printOddLine(m, i == n)
		} else {
			printEven(index, m)
			index++
		}
	}
}

func printOddLine(m int, flg bool) {
	for i := 0; i < m; i++ {
		fmt.Print("#")
	}
	if !flg {
		fmt.Println()
	}
}

func printEven(index, m int) {
	if index%2 == 1 {
		fmt.Print("#")
	}
	for i := 1; i < m; i++ {
		fmt.Print(".")
	}
	if index%2 == 0 {
		fmt.Print("#")
	}
	fmt.Println()
}
