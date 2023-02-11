package main

import "testing"

func runSample(t *testing.T, k int, traps []int, expect int) {
	res := int(solve(k, traps))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 4
	traps := []int{8, 7, 1, 4}
	expect := 0
	runSample(t, k, traps, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	traps := []int{5, 10, 11, 5}
	expect := 21
	runSample(t, k, traps, expect)
}

func TestSample3(t *testing.T) {
	k := 5
	traps := []int{8, 2, 5, 15, 11, 2, 8}
	expect := 9
	runSample(t, k, traps, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	traps := []int{1, 2, 3, 4, 5, 6}
	expect := 6
	runSample(t, k, traps, expect)
}

func TestSample5(t *testing.T) {
	k := 1
	traps := []int{7}
	expect := 0
	runSample(t, k, traps, expect)
}
