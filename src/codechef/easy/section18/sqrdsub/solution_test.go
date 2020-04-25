package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	runSample(t, A, 2)
}

func TestSample2(t *testing.T) {
	A := []int{2, 5, 6}
	runSample(t, A, 2)
}

func TestSample3(t *testing.T) {
	A := []int{2, 5, 4}
	runSample(t, A, 2)
}

func TestSample4(t *testing.T) {
	A := []int{2, 5, 8}
	expect := find(A)
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := make([]int, 10)
	for i := 0; i < 10; i++ {
		A[i] = i + 1
	}
	expect := find(A)
	runSample(t, A, expect)
}

func find(A []int) int64 {
	var res int64

	for i := 0; i < len(A); i++ {
		p := int64(1)
		for j := i; j < len(A); j++ {
			p *= int64(A[j])
			if p%2 == 1 || p%4 == 0 {
				res++
			}
		}
	}
	return res
}
