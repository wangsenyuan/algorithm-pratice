package p3325

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := numberOfSubstrings(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abacb"
	k := 2
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abcde"
	k := 1
	expect := 15
	runSample(t, s, k, expect)
}
