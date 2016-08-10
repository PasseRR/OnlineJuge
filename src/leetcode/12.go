package main

import (
	"bytes"
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var buffer bytes.Buffer
	romans := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I",
	}
	nums := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}
	for i := 0; num > 0; i++ {
		for num >= nums[i] {
			buffer.WriteString(romans[i])
			num -= nums[i]
		}
	}

	return buffer.String()
}

func main() {
	fmt.Printf("%v\n", strings.Repeat("1", 0))
	fmt.Printf("%v\n", intToRoman(9))
}
