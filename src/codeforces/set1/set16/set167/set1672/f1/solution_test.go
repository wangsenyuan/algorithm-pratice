package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	a := countSadness(expect, A)
	b := countSadness(res, A)

	if a != b {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	sort.Ints(res)
	sort.Ints(A)

	if !reflect.DeepEqual(res, A) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func countSadness(arr []int, expect []int) int {
	// 如果arr[i] == expect[i]
	var res int
	for i := 0; i < len(arr); i++ {
		if arr[i] == expect[i] {
			continue
		}
		j := i + 1
		for j < len(arr) && expect[i] != arr[j] {
			j++
		}
		if j == len(arr) {
			return -1
		}
		res++
		arr[i], arr[j] = arr[j], arr[i]
	}
	return res
}

func TestSample1(t *testing.T) {
	A := []int{2, 1}
	expect := []int{1, 2}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 3}
	expect := []int{3, 3, 2, 1}
	runSample(t, A, expect)
}
