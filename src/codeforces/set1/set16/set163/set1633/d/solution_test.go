package main

import "testing"

func runSample(t *testing.T, k int, B []int, C []int, expect int) {
	res := solve(k, B, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 4
	B := []int{1, 7, 5, 2}
	C := []int{2, 6, 5, 2}
	expect := 9
	runSample(t, k, B, C, expect)
}

func TestSample2(t *testing.T) {
	k := 9
	B := []int{5, 2, 5, 6, 3}
	C := []int{5, 9, 1, 9, 7}
	expect := 30
	runSample(t, k, B, C, expect)
}

func TestSample3(t *testing.T) {
	k := 14
	B := []int{11, 4, 6, 2, 8, 16}
	C := []int{43, 45, 9, 41, 15, 38}
	expect := 167
	runSample(t, k, B, C, expect)
}
