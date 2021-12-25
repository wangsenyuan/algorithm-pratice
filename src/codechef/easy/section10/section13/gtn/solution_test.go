package main

import "testing"

func runSample(t *testing.T, N int, X int, arr []int, expect int) {
	query := func(i int) int {
		if i > N {
			return 0
		}
		j := i & (-i)
		// j is cnt
		var s int
		for k := 0; k < j && i+k <= len(arr); k++ {
			s += arr[i+k-1]
		}
		return s
	}

	res := solve(N, X, query)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, X, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, X := 7, 8
	arr := []int{1, 2, 8, 9, 10, 11, 12}
	expect := 3
	runSample(t, N, X, arr, expect)
}

func TestSample2(t *testing.T) {
	N := 8
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	runSample(t, N, 8, arr, 8)

	for i := 0; i < N; i++ {
		X := arr[i]
		expect := i + 1
		runSample(t, N, X, arr, expect)
	}
	runSample(t, N, 9, arr, -1)

}

func TestSample3(t *testing.T) {
	N, X := 6, 13
	arr := []int{1, 2, 8, 10, 11, 12}
	expect := -1
	runSample(t, N, X, arr, expect)
}
