package main
import "fmt"

func main(){
	var n int64
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n/2)
	}else {
		fmt.Println(0-(n+1)/2)
	}
}
