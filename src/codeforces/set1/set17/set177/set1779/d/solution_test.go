package main

import "testing"

func runSample(t *testing.T, a []int, b []int, x []int, expect bool) {
	res := solve(a, b, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 3, 3}
	b := []int{2, 1, 2}
	x := []int{1, 2}
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 4, 6, 3, 4}
	b := []int{3, 1, 2, 3, 2, 3}
	x := []int{3, 2, 3}
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1, 1}
	b := []int{1, 1, 2}
	x := []int{4, 2, 4, 3, 1, 5, 6, 3, 5, 6, 2, 1}
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample5(t *testing.T) {
	a := []int{7, 9, 4, 5, 3, 3, 3, 6, 8, 10, 3, 2, 5}
	b := []int{5, 3, 1, 5, 3, 2, 2, 5, 8, 5, 1, 1, 5}
	x := []int{1, 5, 3, 5, 4, 2, 3, 1}
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample6(t *testing.T) {
	a := []int{7, 9, 4, 5, 3, 3, 3, 6, 8, 10, 3, 2, 5}
	b := []int{5, 3, 1, 5, 3, 2, 2, 5, 8, 5, 1, 1, 5}
	x := []int{1, 5, 3, 4, 2, 3, 1}
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample7(t *testing.T) {
	a := []int{19747843, 2736467, 938578397}
	b := []int{2039844, 2039844, 2039844}
	x := []int{2039844}
	expect := true
	runSample(t, a, b, x, expect)
}
