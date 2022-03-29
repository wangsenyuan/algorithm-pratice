package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if expect {
		var sum int64
		for i := 0; i < len(A)-1; i++ {
			sum += int64(A[i]) * int64(A[i+1])
		}
		if sum&1 != 1 {
			t.Errorf("Sample result %v not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 10}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{7, 11, 145}
	expect := false
	runSample(t, A, expect)
}
