package main

import "testing"

func runSample(t *testing.T, a []int, b []int, x int, expect bool) {
	res := solve(x, a, b)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] > res[i] {
			cnt++
		}
	}

	if cnt != x {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1}
	b := []int{2}
	x := 0
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1}
	b := []int{2}
	x := 1
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{4, 1, 2}
	x := 0
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{4, 1, 2}
	x := 1
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{4, 1, 2}
	x := 1
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{4, 1, 2}
	x := 2
	expect := true
	runSample(t, a, b, x, expect)
}

func TestSample7(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{4, 1, 2}
	x := 3
	expect := false
	runSample(t, a, b, x, expect)
}

func TestSample8(t *testing.T) {
	a := []int{6, 4, 5, 6, 2}
	b := []int{9, 7, 9, 1, 1}
	x := 2
	expect := true
	runSample(t, a, b, x, expect)
}
