package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 2, 3, 6}
	var expect int64 = 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 1}
	var expect int64 = -1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1}
	var expect int64 = 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 1, 1, 2}
	var expect int64 = -1
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 3, 1, 3, 4, 5, 1}
	var expect int64 = 10
	runSample(t, A, expect)
}

// 1 5 3 1
func TestSample6(t *testing.T) {
	A := []int{1, 5, 3, 1}
	// 1, 3, 4, 1
	// 1, 4, 2, 1
	// 1, 2, 2, 1
	// 1, 0, 2, 1
	// 1, 0, 0, 1
	var expect int64 = 5
	runSample(t, A, expect)
}
