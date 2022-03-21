package p2207

import "testing"

func runSample(t *testing.T, text, pattern string, expect int) {
	res := int(maximumSubsequenceCount(text, pattern))

	if res != expect {
		t.Errorf("Sample exect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	text := "abdcdbc"
	pattern := "ac"
	expect := 4
	runSample(t, text, pattern, expect)
}

func TestSample2(t *testing.T) {
	text := "aabb"
	pattern := "ab"
	expect := 6
	runSample(t, text, pattern, expect)
}
