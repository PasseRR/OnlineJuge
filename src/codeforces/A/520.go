package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n        int
		chars    string
		word     string
		alphabet = make(map[string]bool)
	)
	for i := 0; i < 26; i++ {
		alphabet[string(i+97)] = true
	}
	fmt.Scanln(&n)
	fmt.Scanln(&chars)
	if n >= 26 {
		for i := 0; i < n; i++ {
			word = strings.ToLower(string(chars[i]))
			_, ok := alphabet[word]
			if ok {
				delete(alphabet, word)
			}
		}
		if len(alphabet) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}
