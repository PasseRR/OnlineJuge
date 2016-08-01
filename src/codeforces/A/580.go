package main
import (
	"fmt"
)

func main() {
	var(
		n, temp, max int
		current, last int
	)
	fmt.Scanln(&n)
	for i := 0; i < n; i ++ {
		fmt.Scanf("%d", &current)
		if current >= last {
			temp ++
		}else {
			if temp > max {
				max = temp
			}
			temp = 1
		}
		last = current
	}

	if temp > max {
		max = temp
	}
	fmt.Println(max)
}
