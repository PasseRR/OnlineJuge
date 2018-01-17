package page1

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	fmt.Println(multiply("5", "12"))
	fmt.Println(multiply("123", "0"))
	fmt.Println(multiply("123456789", "987654321"))
}
