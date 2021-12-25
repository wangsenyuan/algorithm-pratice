package main

import "testing"

func runSample(t *testing.T, A, B int, expect int) {
	res := solve(A, B)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4, 1)
}
