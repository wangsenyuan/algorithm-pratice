package p2513

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countAnagrams(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "too hot"
	expect := 18
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aa"
	expect := 1
	runSample(t, s, expect)
}
