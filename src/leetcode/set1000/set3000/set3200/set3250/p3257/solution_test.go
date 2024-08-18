package p3257

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := countKConstraintSubstrings(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10101"
	k := 1
	expect := 12
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1010101"
	k := 2
	expect := 25
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "11111"
	k := 1
	expect := 15
	runSample(t, s, k, expect)
}
