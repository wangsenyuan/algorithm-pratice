package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if int(res) != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{1, 2, 3, 4}, 4)
}
