package main

import "testing"

func runSample(t *testing.T, n int, m int, atts []string, expect int) {
	res := solve(n, m, atts)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 2
	atts := []string{
		"000000 2",
		"010100 4",
	}
	expect := 6
	runSample(t, n, m, atts, expect)
}

func TestSample2(t *testing.T) {
	n, m := 6, 3
	atts := []string{
		"000000 2",
		"010100 4",
		"111100 0",
	}
	expect := 0
	runSample(t, n, m, atts, expect)
}

func TestSample3(t *testing.T) {
	n, m := 6, 3
	atts := []string{
		"000000 2",
		"010100 4",
		"111100 2",
	}
	expect := 1
	runSample(t, n, m, atts, expect)
}
