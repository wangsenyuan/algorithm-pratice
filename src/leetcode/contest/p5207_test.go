package contest

import "testing"

func runEqualSubstring(test *testing.T, s, t string, maxCost int, expect int) {
	res := equalSubstring(s, t, maxCost)
	if res != expect {
		test.Errorf("Sample %s %s %d, expect %d, but got %d", s, t, maxCost, expect, res)
	}
}

func TestES1(t *testing.T) {
	runEqualSubstring(t, "abcd", "bcdf", 3, 3)
}

func TestES2(t *testing.T) {
	runEqualSubstring(t, "abcd", "cdef", 3, 1)
}

func TestES3(t *testing.T) {
	runEqualSubstring(t, "abcd", "acde", 0, 1)
}
