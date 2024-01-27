package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, a []int) {
	ask := func(l int, r int) int {
		return countInversions(a[l-1 : r])
	}

	res := solve(len(a), ask)

	res--
	if a[res] != len(a) {
		t.Fatalf("Sample result %d, not correct", res)
	}
}

func countInversions(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				res++
			}
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 2, 4}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 2, 4, 5, 6, 9, 8, 7}
	runSample(t, a)
}

func TestSample3(t *testing.T) {
	n := 100
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	for i := 0; i < 10; i++ {
		rand.Shuffle(n, func(j, k int) {
			arr[j], arr[k] = arr[k], arr[j]
		})
		runSample(t, arr)
	}
}
