package main

import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)
	if x%5 == 0 {
		fmt.Println(x / 5)
	} else {
		fmt.Println(x/5 + 1)
	}
}
