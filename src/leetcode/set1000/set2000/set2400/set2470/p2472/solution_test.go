package p2472

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := maxPalindromes(s, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abaccdbbd"
	k := 3
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "adbcda"
	k := 2
	expect := 0
	runSample(t, s, k, expect)
}
