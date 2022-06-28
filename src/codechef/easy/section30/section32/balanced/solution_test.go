package main

import "testing"

func runSample(t *testing.T, n int, S, T []int, expect string) {
	res := solve(n, S, T)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []int{1, 0, 2}
	T := []int{2, 1, 0}
	expect := "ABA"
	runSample(t, n, S, T, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := []int{0, 2, 1, 3}
	T := []int{1, 2, 3, 0}
	expect := "AABB"
	runSample(t, n, S, T, expect)
}
