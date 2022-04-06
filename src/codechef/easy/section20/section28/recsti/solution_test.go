package main

import "testing"

func runSample(t *testing.T, L []int, expect int) {
	res := solve(L)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := []int{1}
	expect := 3
	runSample(t, L, expect)
}

func TestSample2(t *testing.T) {
	L := []int{2, 5}
	expect := 2
	runSample(t, L, expect)
}

func TestSample3(t *testing.T) {
	L := []int{2, 2, 3, 3}
	expect := 0
	runSample(t, L, expect)
}

func TestSample4(t *testing.T) {
	L := []int{1, 3, 5, 7, 1, 7, 5}
	expect := 1
	runSample(t, L, expect)
}
