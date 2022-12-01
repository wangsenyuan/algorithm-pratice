package main

import (
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := int(solve(A, B))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	B := []int{4, 6}
	expect := 8
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 8, 9, 11}
	B := []int{25, 27, 100, 45}
	expect := 2960
	runSample(t, A, B, expect)
}
