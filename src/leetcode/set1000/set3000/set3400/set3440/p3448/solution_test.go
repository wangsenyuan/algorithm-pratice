package p3448

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countSubstrings(s)
	if res != int64(expect) {
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "12936"
	expect := 11
	runSample(t, s, expect)
}
