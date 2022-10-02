package p2430

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := deleteString(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcabcdabc"
	expect := 2
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := "aaabaab"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaaa"
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "aabaab"
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "aab"
	expect := 2
	runSample(t, s, expect)
}
