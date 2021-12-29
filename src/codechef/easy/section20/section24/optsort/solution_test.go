package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 1, 4}
	expect := 9
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 5, 6, 4, 3, 7, 9, 8}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 5, 3, 8, 7, 9}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 3, 2, 6, 4, 4, 7, 8, 6}
	expect := 5
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 3, 2, 6, 4, 4, 6, 8, 7}
	expect := 4
	runSample(t, A, expect)
}


func TestSample6(t *testing.T) {
	A := []int{5, 1, 3, 4}
	expect := 4
	runSample(t, A, expect)
}
