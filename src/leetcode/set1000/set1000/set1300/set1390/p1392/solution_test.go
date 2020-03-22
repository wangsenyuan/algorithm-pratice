package p1392

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := longestPrefix(s)

	if res != expect {
		t.Errorf("Smaple %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "level", "l")
}

func TestSample2(t *testing.T) {
	runSample(t, "ababab", "abab")
}

func TestSample3(t *testing.T) {
	runSample(t, "leetcodeleet", "leet")
}

func TestSample4(t *testing.T) {
	runSample(t, "a", "")
}
