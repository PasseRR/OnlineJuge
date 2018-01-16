package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 4; i <= n/2; i++ {
		if isPrime(i) && isPrime(n-i) {
			fmt.Printf("%d %d\n", i, n-i)
			break
		}
	}
}

func isPrime(n int) (flg bool) {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flg = true
			return
		}
	}
	return
}
