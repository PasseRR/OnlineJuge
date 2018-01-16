package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		data string
		set  = make(map[string]int)
	)
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	data = string(bytes)
	data = strings.Replace(data, "{", "", 1)
	data = strings.Replace(data, "}", "", 1)
	for _, letter := range strings.Split(data, ", ") {
		if "" != letter {
			_, ok := set[letter]
			if !ok {
				set[letter] = 1
			}
		}
	}

	fmt.Println(len(set))
}
