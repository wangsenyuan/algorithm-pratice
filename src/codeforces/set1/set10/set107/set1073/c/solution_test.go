package main

import "testing"

func runSample(t *testing.T, seq string, x int, y int, expect int) {
	res := solve(seq, x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	seq := "RURUU"
	x, y := -2, 3
	expect := 3
	runSample(t, seq, x, y, expect)
}

func TestSample2(t *testing.T) {
	seq := "RULR"
	x, y := 1, 1
	expect := 0
	runSample(t, seq, x, y, expect)
}

func TestSample3(t *testing.T) {
	seq := "UDUDUD"
	x, y := 0, 1
	expect := -1
	runSample(t, seq, x, y, expect)
}

func TestSample4(t *testing.T) {
	seq := "LRU"
	x, y := 0, -1
	expect := 1
	runSample(t, seq, x, y, expect)
}

func TestSample5(t *testing.T) {
	seq := "DRRD"
	x, y := 1, 3
	expect := 4
	runSample(t, seq, x, y, expect)
}
