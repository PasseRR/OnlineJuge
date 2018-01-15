package page1

import "testing"

func TestConvert(t *testing.T) {
	t.Logf("%v", convert("ABCD", 3))
}
