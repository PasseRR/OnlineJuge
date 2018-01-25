package page1

import "testing"

func TestIsMatch2(t *testing.T) {
	//t.Logf("%v\n", isMatch2("ho", "ho**"))
	//t.Logf("%v\n", isMatch2("aa", "*"))
	//t.Logf("%v\n", isMatch2("a", "aa"))
	//t.Logf("%v\n", isMatch2("aa", "a"))
	//t.Logf("%v\n", isMatch2("aa", "aa"))
	//t.Logf("%v\n", isMatch2("aaa", "aa"))
	//t.Logf("%v\n", isMatch2("aa", "a*"))
	//t.Logf("%v\n", isMatch2("ab", "?*"))
	//t.Logf("%v\n", isMatch2("aab", "c*a*b*"))
	t.Logf("%v\n", isMatch2("abefcdgiescdfimde", "ab*cd?i*de"))
	//t.Logf("%v\n", isMatch2("zacabz", "*a?b*"))
	//t.Logf("%v\n", isMatch2("", "*"))
	//t.Logf("%v\n", isMatch2("", "?"))
}
