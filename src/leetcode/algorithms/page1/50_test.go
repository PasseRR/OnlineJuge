package page1

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(34.00515, -3))
	fmt.Println(myPow(0.00001, 2147483647))
}
