package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1337}
	var expect = 446
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 8, 10}
	var expect = 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	var expect = 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{7, 77, 777}
	var expect = 260
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 3, 2}
	var expect = 2
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{4, 3, 2}
	var expect = 3
	runSample(t, A, expect)
}

func TestSample7(t *testing.T) {
	A := []int{4, 5}
	// 2, 2, 1
	var expect = 3
	runSample(t, A, expect)
}

func TestSample8(t *testing.T) {
	A := []int{4, 2}
	// 2, 2, 1
	var expect = 2
	runSample(t, A, expect)
}

func TestSample9(t *testing.T) {
	A := []int{1, 5}
	// 1, 1, 3
	var expect = 3
	runSample(t, A, expect)
}
