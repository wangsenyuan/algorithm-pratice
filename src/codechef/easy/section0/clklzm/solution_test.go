package main

import "testing"

func runSample(t *testing.T, n, m int, Z []int, L []int, R []int, K []int, expect int) {
	res := solve(n, m, Z, L, R, K)
	if res != expect {
		t.Errorf("sample %d %d %v %v %v %v, expect %d, but got %d", n, m, Z, L, R, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 5
	Z := []int{1, 2, 3, 4, 5}
	L := []int{1, 1, 2, 3, 4}
	R := []int{2, 5, 4, 5, 5}
	K := []int{3, 1, 3, 3, 3}
	expect := 6
	runSample(t, n, m, Z, L, R, K, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 2
	Z := []int{2, 5, 2}
	L := []int{1, 2}
	R := []int{2, 3}
	K := []int{2, 2}
	expect := -1
	runSample(t, n, m, Z, L, R, K, expect)
}
