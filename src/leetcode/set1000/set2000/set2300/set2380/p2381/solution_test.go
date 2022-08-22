package p2381

import "testing"

func runSample(t *testing.T, s string, shifts [][]int, expect string) {
	res := shiftingLetters(s, shifts)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abc"
	shifts := [][]int{
		{0, 1, 0}, {1, 2, 1}, {0, 2, 1},
	}
	expect := "ace"
	runSample(t, s, shifts, expect)
}

func TestSample2(t *testing.T) {
	s := "dztz"
	shifts := [][]int{
		{0, 0, 0}, {1, 1, 1},
	}
	expect := "catz"
	runSample(t, s, shifts, expect)
}
