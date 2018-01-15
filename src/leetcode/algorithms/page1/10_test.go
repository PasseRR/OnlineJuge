package page1

import "testing"

func TestIsMatch(t *testing.T) {
	t.Logf("%v\n", isMatch("a", "ab*"))
}
