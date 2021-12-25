package main

import "testing"

func runSample(t *testing.T, S, T string, nums []int, expect []int64) {
	solver := NewSolver(S, T)

	for i, num := range nums {
		res := solver.Query(num)

		if res != expect[i] {
			t.Errorf("Sample fail at %d, expect %d, but got %d", i, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	S := "codechef"
	T := "chefcode"
	nums := []int{4, 12, 1455}
	expect := []int64{0, 1, 181}
	runSample(t, S, T, nums, expect)
}
