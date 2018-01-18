package page1

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	words, result := map[string][]string{}, [][]string{}

	for _, word := range strs {
		sortedWord := sortString(word)
		words[sortedWord] = append(words[sortedWord], word)
	}

	for _, value := range words {
		result = append(result, value)
	}

	return result
}
