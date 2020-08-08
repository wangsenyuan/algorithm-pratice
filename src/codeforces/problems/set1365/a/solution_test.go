package main

import "testing"

func runSample(t *testing.T, n, m int, A [][]int, expect string) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{0, 0},
		{0, 0},
	}
	expect := "Vivek"
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	A := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}
	expect := "Ashish"
	runSample(t, n, m, A, expect)
}
