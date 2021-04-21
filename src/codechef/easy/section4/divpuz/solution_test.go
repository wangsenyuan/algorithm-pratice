package main

import "testing"

func runSample(t *testing.T, A []int) {
	B := solve(A)

	n := len(A)
	for i := 0; i < len(A); i++ {
		if B[i]%int64(A[i]) != 0 {
			t.Fatalf("Sample result %v, not correct", B)
		}
		var tmp int64
		if i == 0 {
			tmp = B[n-1] * B[1]
		} else if i == n-1 {
			tmp = B[0] * B[n-2]
		} else {
			tmp = B[i-1] * B[i+1]
		}

		if tmp%B[i] != 0 {
			t.Fatalf("Sample result %v, not correct", B)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 9}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := []int{6, 10, 15}
	runSample(t, A)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, 4, 5}
	runSample(t, A)
}
