package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {

	res := solve(a)
	x := operate(a, expect)

	y := operate(a, res)

	if x != y {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func operate(arr []int, x []int) int {
	n := len(arr)
	a := make([]int, n)
	copy(a, arr)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[x[i]-1] += i + 1
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		a[i] += diff[i]
	}
	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[j] < a[i] {
				res++
			}
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	expect := []int{1, 1, 1, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 2, 4, 5}
	expect := []int{1, 4, 3, 2, 1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3, 1}
	expect := []int{1, 3, 3}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 2, 1}
	expect := []int{3, 2, 3}
	runSample(t, a, expect)
}
