package main
import "fmt"

func main() {
	var (
		n, num, lastOdd, lastEven, oddNum, evenNum int
	)
	fmt.Scanln(&n)
	for i := 1; i <= n; i ++ {
		fmt.Scanf("%d", &num)
		if num%2 == 1 {
			lastOdd = i
			oddNum ++
		}else {
			lastEven = i
			evenNum ++
		}

		if oddNum > 0 && evenNum > 0 {
			if oddNum > evenNum {
				fmt.Println(lastEven)
				break;
			}

			if evenNum > oddNum {
				fmt.Println(lastOdd)
				break;
			}
		}
	}
}
