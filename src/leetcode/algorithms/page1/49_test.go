package page1

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
