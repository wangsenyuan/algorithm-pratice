package main

import "testing"

func runSample(t *testing.T, n int, K int, A []int, expect int, inc [][]int, desc [][]int) {
	res := solve(n, A, K, inc, desc)

	if res != expect {
		t.Errorf("sample %d %d %v %v %v, expect %d but got %d", n, K, A, inc, desc, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 10
	A := []int{2, 3, 5, 4}
	inc := [][]int{{1, 2}}
	desc := [][]int{{3, 4}}
	expect := 1
	runSample(t, n, k, A, expect, inc, desc)
}

func TestSample2(t *testing.T) {
	n, k := 5, 10
	A := []int{-1, -1, -1, -1, -1}
	inc := [][]int{{1, 3}}
	desc := [][]int{{3, 5}}
	expect := 8
	runSample(t, n, k, A, expect, inc, desc)
}
