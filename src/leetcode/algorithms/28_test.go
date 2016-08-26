package algorithms

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	fmt.Printf("%v\n", index("aaab", "ab"))
	fmt.Printf("%v\n", strStr("aaab", "ab"))
}
