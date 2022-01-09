package p2135

import "testing"

func runSample(t *testing.T, start []string, target []string, expect int) {
	res := wordCount(start, target)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := []string{"ant", "act", "tack"}
	target := []string{"tack", "act", "acti"}
	expect := 2
	runSample(t, start, target, expect)
}
