package main

import "testing"

func runSample(t *testing.T, n int, A []int) {
	begin, end := solve(n, A)
	if begin < 0 {
		t.Errorf("The answer should always exists")
		return
	}

	var sum int
	for i := begin; i <= end; i++ {
		sum += A[i]
	}
	if sum%n != 0 {
		t.Errorf("sample %d %v, answer %d, %d, not correct", n, A, begin, end)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{4, 6, 10}
	runSample(t, n, A)
}
