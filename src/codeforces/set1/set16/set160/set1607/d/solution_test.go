package main

import "testing"

func runSample(t *testing.T, A []int, color string, expect bool) {
	res := solve(A, color)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 5, 2}
	color := "BRBR"
	expect := true
	runSample(t, A, color, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1}
	color := "BB"
	expect := false
	runSample(t, A, color, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 1, 4, 2, 5}
	color := "RBRRB"
	expect := true
	runSample(t, A, color, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 1, 5, 1, 5}
	color := "RBRRB"
	expect := false
	runSample(t, A, color, expect)
}
