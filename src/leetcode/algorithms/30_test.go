package algorithms

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	fmt.Printf("%v\n", findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Printf("%v\n", findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Printf("%v\n", findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Printf("%v\n", findSubstring("a", []string{"a"}))
	fmt.Printf("%v\n", findSubstring("aaa", []string{"aa", "aa"}))
	fmt.Printf("%v\n", findSubstring("aaaaaaaa", []string{"aa", "aa", "aa"}))
}
