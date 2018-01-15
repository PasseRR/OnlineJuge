package page1

import "fmt"

func countAndSay(n int) string {
	result := "1"
	if n > 1 {
		for i := 0; i < n-1; i++ {
			count, current, temp := 1, "", result[0:1]
			for j := 1; j < len(result); j++ {
				b := result[j : j+1]
				if temp == b {
					count++
				} else {
					current += fmt.Sprintf("%d%s", count, temp)
					count = 1
					temp = b
				}
			}
			result = current + fmt.Sprintf("%d%s", count, temp)
		}
	}

	return result
}
