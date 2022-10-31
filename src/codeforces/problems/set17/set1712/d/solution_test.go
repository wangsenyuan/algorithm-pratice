package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{2, 4, 1}
	expect := 4
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{1, 9, 84}
	expect := 168
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	A := []int{10, 2, 6}
	expect := 10
	runSample(t, k, A, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	A := []int{179, 17, 1000000000}
	expect := 1000000000
	runSample(t, k, A, expect)
}
