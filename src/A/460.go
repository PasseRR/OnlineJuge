package main
import "fmt"

func main(){
	var(
		n, m int
		total int
	)
	fmt.Scanf("%d %d\n", &n, &m)
	total += n
	for i := n-m; i >= 0; i -= m {
		i ++
		total ++
	}

	fmt.Println(total)
}
