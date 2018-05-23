package main

import "testing"

func runSample(t *testing.T, n int, s string, nums []int, expect int64) {
	res := solve(n, []byte(s), nums)

	if res != expect {
		t.Errorf("Sample %d %s %v, expect %d, but got %d", n, s, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	s := "()()"
	nums := []int{-1, -2, 3, 4}
	runSample(t, n, s, nums, 7)
}

func TestSample2(t *testing.T) {
	n := 4
	s := "(()]"
	nums := []int{-1, -2, 3, 4}
	runSample(t, n, s, nums, 1)
}

func TestSample3(t *testing.T) {
	n := 4
	s := "[{]{"
	nums := []int{1, 2, 3, 4}
	runSample(t, n, s, nums, 0)
}

func TestSample4(t *testing.T) {
	n := 6
	s := "((()))"
	nums := []int{1, 2, 3, 4, 5, 6}
	runSample(t, n, s, nums, 21)
}

func TestSample5(t *testing.T) {
	n := 6
	s := "((())]"
	nums := []int{1, 2, 3, 4, 5, 6}
	runSample(t, n, s, nums, 14)
}
