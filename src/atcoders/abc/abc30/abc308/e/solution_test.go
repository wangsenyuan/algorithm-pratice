package main

import "testing"

func runSample(t *testing.T, A []int, S string, expect int) {
	res := solve(A, S)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 0, 2}
	S := "MEEX"
	expect := 3
	runSample(t, A, S, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 2, 0, 0, 2, 0, 2, 0, 0, 0, 0, 0, 2, 2}
	S := "EXMMXXXEMEXEXMM"
	expect := 13
	runSample(t, A, S, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 0, 2, 2}
	S := "MEEXX"
	expect := 6
	runSample(t, A, S, expect)
}
