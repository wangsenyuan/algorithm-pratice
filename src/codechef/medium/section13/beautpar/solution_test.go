package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect bool) {
	res := solve(n, k, A)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var pos int

	for i := 0; i < k; i++ {
		var sum int64
		var x int
		for x < res[i] && pos < n {
			sum += int64(A[pos])
			x++
		}
		if x < res[i] || sum <= 0 {
			t.Fatalf("result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	k := 2
	A := []int{2, -1, 1, 3}
	expect := true
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	k := 2
	A := []int{2, -1}
	expect := false
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	k := 1
	A := []int{2, -1}
	expect := true
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n := 1
	k := 1
	A := []int{1}
	expect := true
	runSample(t, n, k, A, expect)
}
