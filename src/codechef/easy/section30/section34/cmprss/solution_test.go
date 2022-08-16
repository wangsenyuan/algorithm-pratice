package main

import "testing"

func runSample(t *testing.T, A []int, expect string) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 5, 6, 8, 9, 10, 11, 12, 15, 17}
	expect := "1...3,5,6,8...12,15,17"
	runSample(t, A, expect)
}
