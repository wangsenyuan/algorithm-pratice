package main

import "testing"

func runSample(t *testing.T, A []int) {
	res := solve(A)

	if len(res)/2 > 20 {
		t.Errorf("Sample result exceeds the limit, got %d steps", len(res)/2)
	} else {
		B := make([]int, len(A))
		copy(B, A)
		for i := 0; i < len(res); i += 2 {
			x := res[i][1]
			for _, j := range res[i+1] {
				B[j-1] += x
			}
		}
		for i := 1; i < len(B); i++ {
			if B[i] <= B[i-1] {
				t.Errorf("Sample result %v, got wrong result %v", res, B)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 4, 4}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := []int{8, 8, 5, 2, 1}
	runSample(t, A)
}
