package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect int) {
	B := make([]int, n)
	copy(B, A)
	cnt, res := solve(n, A)
	if (expect < 0) != (cnt < 0) {
		t.Errorf("Sample %v, expect %d, but got %d", B, expect, cnt)
		return
	}
	if expect < 0 {
		return
	}

	if cnt > n*n {
		t.Errorf("Sample %v, expect %d, but got %d", B, expect, cnt)
		return
	}
	for i := 0; i < cnt; i++ {
		j := res[i] - 1
		cycle(&B[j], &B[j+1], &B[j+2])
	}

	if !sort.IntsAreSorted(B) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{5, 4, 3, 2, 1}
	expect := 6
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	A := []int{8, 4, 5, 2, 3, 6, 7, 3}
	expect := 13
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	A := []int{5, 2, 1, 6, 4, 7, 3}
	expect := -1
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 6
	A := []int{1, 2, 3, 3, 6, 4}
	expect := 4
	runSample(t, n, A, expect)
}
