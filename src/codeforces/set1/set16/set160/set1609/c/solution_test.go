package main

import "testing"

func runSample(t *testing.T, e int, A []int, expect int64) {
	res := solve(e, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	e := 3
	A := []int{10, 2, 1, 3, 1, 19, 3}
	var expect int64 = 2
	runSample(t, e, A, expect)
}

func TestSample2(t *testing.T) {
	e := 2
	A := []int{1, 13, 1}
	var expect int64
	runSample(t, e, A, expect)
}

func TestSample3(t *testing.T) {
	e := 3
	A := []int{2, 4, 2, 1, 1, 1, 1, 4, 2}
	var expect int64 = 4
	runSample(t, e, A, expect)
}

func TestSample5(t *testing.T) {
	e := 1
	A := []int{1, 2, 1, 1}
	var expect int64 = 5
	runSample(t, e, A, expect)
}

func TestSample6(t *testing.T) {
	e := 2
	A := []int{1, 1, 1, 422549, 1, 1, 880667, 81267, 1, 1}
	var expect int64 = 10
	runSample(t, e, A, expect)
}

func TestSample7(t *testing.T) {
	e := 2
	A := []int{1, 7, 1, 1, 1, 1, 3, 1, 4, 1}
	var expect int64 = 7
	runSample(t, e, A, expect)
}
