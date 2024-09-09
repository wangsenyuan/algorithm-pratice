package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	x := sum(expect)
	y := sum(res)

	if x != y {
		t.Fatalf("")
	}
}

func sum(arr []int) int {
	var res int
	for _, num := range arr {
		res += num
	}
	return res
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 2, 1}
	expect := []int{1, 2, 3, 2, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 6, 8}
	expect := []int{10, 6, 6}
	runSample(t, a, expect)
}
