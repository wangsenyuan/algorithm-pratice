package p2483

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countPalindromes(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "103301"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0000000"
	expect := 21
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "9999900000"
	expect := 2
	runSample(t, s, expect)
}