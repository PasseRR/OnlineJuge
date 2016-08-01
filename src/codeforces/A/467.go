package main
import "fmt"

func main(){
	var (
		n, p, q int
		count int
	)
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &p, &q)
		if q - p >= 2 {
			count += 1
		}
	}

	fmt.Println(count)
}