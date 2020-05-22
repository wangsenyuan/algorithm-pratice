package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, []int{42}, 147483634)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, []int{5, 4, 3, 2, 1}, 986095186)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, []int{1194533513, 122420337, 1448417648, 120078455}, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, []int{31, 2047, 2147483647}, 468598063)
}
