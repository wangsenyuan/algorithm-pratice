package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []int{2, 7}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, []int{2, 2, 3, 4}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{2, 2, 4}, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, []int{2, 3, 5, 7}, 4)
}
