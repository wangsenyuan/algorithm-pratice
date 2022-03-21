package main

import "testing"

func runSample(t *testing.T, A, B []int, expect1, expect2 int) {
	ans1, ans2 := solve(A, B)

	if ans1 != expect1 || ans2 != expect2 {
		t.Errorf("Sample expect %d, %d, but got %d %d", expect1, expect2, ans1, ans2)
	}
}

/**
*
2 1 1 3 5
1 2 2 4 5
**/
func TestSample1(t *testing.T) {
	A := []int{2, 1, 1, 3, 5}
	B := []int{1, 2, 2, 4, 5}
	expect1 := 2
	expect2 := 8
	runSample(t, A, B, expect1, expect2)
}
