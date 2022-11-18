package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 5, 7}
	K := 4
	expect := 6
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	K := 4
	expect := 20
	runSample(t, A, K, expect)
}


func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	K := 3
	expect := 14
	runSample(t, A, K, expect)
}
