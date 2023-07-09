package p2767

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minimumBeautifulSubstrings(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1011"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111"
	expect := 3
	runSample(t, s, expect)
}
