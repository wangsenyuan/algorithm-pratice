package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	x := f(expect)
	y := f(res)

	if x != y {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func f(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	cur := (arr[0] | arr[1]) - arr[1]

	for i := 2; i < len(arr); i++ {
		cur = (cur | arr[i]) - arr[i]
	}
	return cur
}

func TestSample1(t *testing.T) {
	a := []int{4, 0, 11, 6}
	expect := []int{11, 6, 4, 0}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{13}
	expect := []int{13}
	runSample(t, a, expect)
}
