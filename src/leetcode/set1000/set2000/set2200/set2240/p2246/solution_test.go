package p2246

import (
	"testing"
)

func runSample(t *testing.T, parent []int, s string, expect int) {
	res := longestPath(parent, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parent := []int{-1, 0, 0, 1, 1, 2}
	s := "abacbe"
	expect := 3
	runSample(t, parent, s, expect)
}

func TestSample2(t *testing.T) {
	parent := []int{-1, 0, 0, 0}
	s := "aabc"
	expect := 3
	runSample(t, parent, s, expect)
}
