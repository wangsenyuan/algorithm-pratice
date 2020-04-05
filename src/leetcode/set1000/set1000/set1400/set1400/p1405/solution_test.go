package p1405

import "testing"

func runSample(t *testing.T, a, b, c int, expect int) {
	res := longestDiverseString(a, b, c)
	if len(res) != expect {
		t.Fatalf("sample %d %d %d, expect %d, but got %s", a, b, c, expect, res)
	}

	if !check(res) {
		t.Fatalf("sample %d %d %d, expect %d, but got %s", a, b, c, expect, res)
	}
}

func check(s string) bool {
	for i := 0; i < len(s); i++ {

		if i >= 2 && s[i-1] == s[i] && s[i-2] == s[i] {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	a, b, c := 1, 1, 7
	expect := len("ccaccbcc")
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a, b, c := 7, 1, 0
	expect := len("aabaa")
	runSample(t, a, b, c, expect)
}
