package main

import "testing"

func runSample(t *testing.T, A []int) {
	R := solve(A)
	n := len(A)

	var cnt int
	for i := 0; i < n; i++ {
		if R[i] != int64(A[i]) {
			cnt++
		}
		if i > 0 {
			if gcd(R[i], R[i-1]) == 1 {
				t.Fatalf("Sample result %v, not correct", R)
			}
		}
	}
	if cnt > (2*n+2)/3 {
		t.Errorf("Sample result change too lot %d", cnt)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 12, 5}
	runSample(t, A)
}
