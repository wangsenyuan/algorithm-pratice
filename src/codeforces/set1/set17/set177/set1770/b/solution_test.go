package main

import "testing"

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	a := getScore(expect, k)
	b := getScore(res, k)

	if a != b {
		t.Fatalf("Sample result %v got wrong score %d, expect %d", res, b, a)
	}
}

func getScore(arr []int, k int) int {
	n := len(arr)
	var res int
	for i := 0; i < n-k; i++ {
		var a, b = n, 0
		for j := 0; j < k; j++ {
			a = min(a, arr[i+j])
			b = max(b, arr[i+j])
		}
		res = max(res, a+b)
	}

	return res
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	expect := []int{5, 1, 2, 3, 4}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 1
	expect := []int{1, 2, 3, 4, 5}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 6, 6
	expect := []int{3, 2, 4, 1, 6, 5}
	runSample(t, n, k, expect)
}
