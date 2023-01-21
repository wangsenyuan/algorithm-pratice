package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int64) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	A := []int{10, 10, 3, 1}
	expect := int64(11)
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{1}
	expect := int64(0)
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 12
	A := []int{10, 10, 29}
	expect := int64(13)
	runSample(t, k, A, expect)
}

func TestSample4(t *testing.T) {
	k := 51
	A := []int{5, 74, 89, 45, 18, 69, 67, 67, 11, 96, 23, 59}
	expect := int64(60)
	runSample(t, k, A, expect)
}
