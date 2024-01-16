package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, k int, expect []int) {
	res := solve(k, a)

	x := getValue(a, expect)
	y := getValue(a, res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func getValue(a []int, res []int) int {
	i, j, x := res[0], res[1], res[2]
	i--
	j--
	return (a[i] ^ x) & (a[j] ^ x)
}

func TestSample1(t *testing.T) {
	a := []int{3, 9, 1, 4, 13}
	k := 4
	expect := []int{1, 3, 14}
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 1}
	k := 1
	expect := []int{1, 3, 0}
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{144, 1580, 1024, 100, 9, 13}
	k := 12
	expect := []int{5, 6, 4082}
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 11, 10, 1, 6, 9, 11, 0, 5}
	k := 4
	expect := []int{2, 7, 4}
	runSample(t, a, k, expect)
}
