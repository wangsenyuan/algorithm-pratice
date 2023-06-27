package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = i
	}

	if len(res) > n+5 {
		t.Fatalf("Sample result took %d steps > %d", len(res), n+5)
	}

	for _, cur := range res {
		x, y := cur[0], cur[1]
		arr[x] = (arr[x] + arr[y] - 1) / arr[y]
	}
	sort.Ints(arr[1:])
	if arr[n] != 2 || arr[n-1] != 1 {
		t.Fatalf("Sample result %v, giving wrong answer %v", res, arr)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3)
}

func TestSamaple2(t *testing.T) {
	for i := 10; i <= 100000; i *= 100 {
		runSample(t, i)
	}
}
