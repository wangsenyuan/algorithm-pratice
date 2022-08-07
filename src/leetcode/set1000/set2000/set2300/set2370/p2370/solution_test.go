package p2370

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := longestIdealString(s, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "acfgbd"
	k := 2
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	k := 3
	expect := 4
	runSample(t, s, k, expect)
}
