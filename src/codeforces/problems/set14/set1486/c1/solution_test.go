package main

import "testing"

func runSample(t *testing.T, arr []int) {
	n := len(arr)

	query := func(l, r int) int {
		first := -1
		second := -l
		for i := l; i <= r; i++ {
			if first < 0 || arr[i-1] > arr[first-1] {
				second = first
				first = i
			} else if second < 0 || arr[i-1] > arr[second-1] {
				second = i
			}
		}
		return second
	}

	res := solve(n, query)
	for i := 0; i < n; i++ {
		if i+1 == res {
			continue
		}
		if arr[i] > arr[res-1] {
			t.Fatalf("Sample %v, result %d not the maximum", arr, res)
		}
	}
}

func TestSample1(t *testing.T) {
	arr := []int{5, 1, 4, 2, 3}
	runSample(t, arr)
}
