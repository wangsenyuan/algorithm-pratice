package main

import "testing"

func runSample(t *testing.T, n int, A []int) {
	res := solve(n, A)
	cnt := make(map[int]int)

	for _, a := range res {
		cnt[a[0]]++
		cnt[a[1]]++
		if (A[a[0]-1]+A[a[1]-1])&1 != 0 {
			t.Errorf("%d %d not divides 2", A[a[0]-1], A[a[1]-1])
		}
	}

	if len(cnt) != 2*n-2 {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3, 4, 5, 6}
	runSample(t, n, A)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{5, 7, 9, 10}
	runSample(t, n, A)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 3, 3, 4, 5, 90, 100, 101, 2, 3}
	runSample(t, n, A)
}
