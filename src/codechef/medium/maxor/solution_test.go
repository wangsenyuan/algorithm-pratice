package main

import "testing"

func runSubTask1(t *testing.T, n int, A []int, expect int64) {
	res := solveTask1(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSubTask1(t, 3, []int{1, 2, 3}, 2)
}

func runSubTask2(t *testing.T, n int, A []int, expect int64) {
	res := solveTask2(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample2(t *testing.T) {
	runSubTask2(t, 3, []int{1, 2, 3}, 2)
}
