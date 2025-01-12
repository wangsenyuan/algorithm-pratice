package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res == expect {
		return
	}
	if expect < 0 || res < 0 {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	var arr []int
	for i := range len(a) {
		if i+1 == res {
			continue
		}
		arr = append(arr, a[i])
	}
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 1; i < len(a); i++ {
		if arr[i] != arr[0]+d*i {
			t.Fatalf("Sample result %d, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 6, 8, 7, 4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 4, 8}
	expect := -1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 8}
	expect := 7
	runSample(t, a, expect)
}
