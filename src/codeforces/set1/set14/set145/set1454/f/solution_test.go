package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(A)
	x, y, z := res[0], res[1], res[2]
	if x+y+z != n {
		t.Fatalf("Sample result %v, not correct, sum get %d != %d", res, x+y+z, n)
	}

	mn := NewTree(1<<30, A, func(a, b int) int {
		return min(a, b)
	})
	mx := NewTree(0, A, func(a, b int) int {
		return max(a, b)
	})
	a := mx.Get(0, x)
	b := mn.Get(x, x+y)
	c := mx.Get(x+y, n)

	if a != b || b != c || c != a {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 3, 3, 4, 4, 3, 4, 2, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 9, 1, 7, 3, 9, 4, 1}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 1, 4, 2, 4, 3, 3, 1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 1, 4, 2, 4, 3, 3, 1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{6, 1, 4, 1, 6, 7, 1, 3, 6, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{2, 4, 2, 2, 8, 4, 8, 2, 6, 2}
	expect := true
	runSample(t, A, expect)
}
