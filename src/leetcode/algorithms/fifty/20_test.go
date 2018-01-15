package fifty

import "testing"

func TestIsValid(t *testing.T) {
	t.Logf("%v", isValid("((()))"))
}
