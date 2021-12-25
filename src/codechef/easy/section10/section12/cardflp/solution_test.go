package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int, cnt int) {
	res, cnt2 := solve(len(A), A, B)

	if res != expect || cnt != cnt2 {
		t.Errorf("Sample expect %d, %d, but got %d %d", expect, cnt, res, cnt2)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 6, 8}
	B := []int{2, 1, 2}
	expect := 2
	cnt := 2
	runSample(t, A, B, expect, cnt)
}

func TestSample2(t *testing.T) {
	A := []int{0, 2, 0}
	B := []int{2, 0, 8}
	expect := 0
	cnt := 0
	runSample(t, A, B, expect, cnt)
}
func TestSample3(t *testing.T) {
	A := []int{1, 2, 1}
	B := []int{2, 3, 6}
	expect := 2
	cnt := 2
	runSample(t, A, B, expect, cnt)
}
