package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{20, 22, 19, 84}
	k := 2
	expect := 2
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{9, 5, 3, 2, 1}
	k := 1
	expect := 3
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{9, 5, 3, 2, 1}
	k := 2
	expect := 2
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := []int{22, 12, 16, 4, 3, 22, 12}
	k := 2
	expect := 3
	runSample(t, A, k, expect)
}

func TestSample5(t *testing.T) {
	A := []int{22, 12, 16, 4, 3, 22, 12}
	k := 3
	expect := 1
	runSample(t, A, k, expect)
}

func TestSample6(t *testing.T) {
	A := []int{3, 9, 12, 3, 9, 12, 3, 9, 12}
	k := 3
	expect := 0
	runSample(t, A, k, expect)
}
