package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect []int) {
	res := solve(a, k)

	x := mul(res)
	y := mul(expect)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func mul(arr []int) int {
	r := 1
	for _, num := range arr {
		r *= num
	}
	return r
}

func TestSample1(t *testing.T) {
	a := []int{4, 11, 8, 2, 1, 13}
	k := 4
	expect := []int{1, 2, 4, 11, 13}

	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 6}
	k := 1
	// expect := []int{1, 2, 4, 11, 13}

	runSample(t, a, k, nil)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3, 1, 5, 3}
	k := 7
	// expect := []int{1, 2, 4, 11, 13}

	runSample(t, a, k, nil)
}

func TestSample4(t *testing.T) {
	a := []int{8, 9, 4, 17, 11, 5}
	k := 3
	expect := []int{9, 11, 17}

	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 2, 2, 2, 2}
	k := 6
	expect := []int{2, 2, 2, 2}

	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{6, 10, 3, 2, 7, 9, 10, 4, 7, 6}
	k := 3
	expect := []int{3, 7, 7, 9}

	runSample(t, a, k, expect)
}
