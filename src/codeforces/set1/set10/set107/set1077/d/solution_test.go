package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, arr []int, k int, expect []int) {
	sort.Ints(arr)
	res := solve(arr, k)

	a := countCopies(arr, expect)
	b := countCopies(arr, res)

	if a != b {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func countCopies(arr []int, expect []int) int {
	f1 := getFreq(expect)
	f2 := getFreq(arr)
	res := len(arr)
	for k, v := range f1 {
		res = min(res, f2[k]/v)
	}

	return res
}

func getFreq(arr []int) map[int]int {
	res := make(map[int]int)
	for _, num := range arr {
		res[num]++
	}

	return res
}

func TestSample1(t *testing.T) {
	k := 4
	arr := []int{1, 3, 1, 3, 10, 3, 7, 7, 12, 3}
	expect := []int{1, 3, 3, 7}
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	arr := []int{2, 2, 2, 3, 1}
	expect := []int{2, 1}
	runSample(t, arr, k, expect)
}
