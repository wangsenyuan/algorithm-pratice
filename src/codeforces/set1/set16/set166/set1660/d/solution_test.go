package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if checkResult(a, expect) != checkResult(a, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func checkResult(a []int, expect []int) int {
	var neg int
	var cnt int
	for i := expect[0]; i < len(a)-expect[1]; i++ {
		if a[i] < 0 {
			neg ^= 1
		}
		cnt += checkValue(abs(a[i]) == 2)
	}

	if neg == 1 {
		return -1
	}

	return cnt
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, -1, 2}
	expect := []int{0, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, -2}
	expect := []int{3, 0}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 0, -2, 2, -1}
	expect := []int{2, 0}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-1, -2, -2}
	expect := []int{1, 0}
	runSample(t, a, expect)
}
