package main

import "testing"

func runSample(t *testing.T, m int64, A []int, expect int) {
	res := solve(m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var m int64 = 4
	A := []int{1, 3, 5, 7}
	expect := 3
	runSample(t, m, A, expect)
}

func TestSample2(t *testing.T) {
	var m int64 = 4
	A := []int{5, 5, 5, 5}
	expect := 4
	runSample(t, m, A, expect)
}

func TestSample3(t *testing.T) {
	var m int64 = 0
	A := []int{4, 6, 9, 12}
	expect := 1
	runSample(t, m, A, expect)
}

func TestSample4(t *testing.T) {
	var m int64 = 10
	A := []int{15, 9, 3, 8, 14, 17}
	expect := 7
	runSample(t, m, A, expect)
}
