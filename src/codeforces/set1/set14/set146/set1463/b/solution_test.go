package main

import "testing"

func runSample(t *testing.T, A []int) {
	B := solve(A)

	n := len(A)

	for i := 0; i+1 < n; i++ {
		if B[i]%B[i+1] == 0 {
			continue
		}
		if B[i+1]%B[i] == 0 {
			continue
		}
		t.Fatalf("Sample result %v, not correct", B)
	}

	var s1, s2 int64

	for i := 0; i < n; i++ {
		s1 += int64(A[i])
		s2 += int64(abs(A[i] - B[i]))
	}
	if 2*s2 > s1 {
		t.Fatalf("Sample result %v, get sum 2 * %d > %d ", B, s2, s1)
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := []int{4, 6}
	runSample(t, A)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1000000000}
	runSample(t, A)
}

func TestSample4(t *testing.T) {
	A := []int{3, 4, 8, 1, 2, 3}
	runSample(t, A)
}
