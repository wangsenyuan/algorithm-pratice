package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {

	x := make([]int, len(a))
	y := make([]int, len(b))
	copy(x, a)
	copy(y, b)

	res, ord := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var sum int

	for _, u := range ord {
		u--
		sum += x[u]
		if y[u] > 0 {
			x[y[u]-1] += x[u]
		}
	}

	if sum != res {
		t.Fatalf("Sample result %v, not getting the expect result %d", ord, expect)
	}

}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, -1}
	expect := 10
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 100}
	b := []int{2, -1}
	expect := 99
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-10, -1, 2, 2, 5, -2, -3, -4, 2, -6}
	b := []int{-1, -1, 2, 2, -1, 5, 5, 7, 7, 9}
	expect := -9
	runSample(t, a, b, expect)
}
