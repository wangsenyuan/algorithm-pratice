package main

import "testing"

func runSample(t *testing.T, s string, x int, y int, expect int) {
	res := solve(s, x, y)
	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0?1"
	x, y := 2, 3
	expect := 4
	runSample(t, s, x, y, expect)
}

func TestSample2(t *testing.T) {
	s := "?????"
	x, y := 13, 37
	expect := 0
	runSample(t, s, x, y, expect)
}

func TestSample3(t *testing.T) {
	s := "?10?"
	x, y := 239, 7
	expect := 28
	runSample(t, s, x, y, expect)
}

func TestSample4(t *testing.T) {
	s := "01101001"
	x, y := 5, 7
	expect := 96
	runSample(t, s, x, y, expect)
}
