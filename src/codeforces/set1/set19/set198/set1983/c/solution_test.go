package main

import "testing"

func runSample(t *testing.T, a []int, b []int, c []int, expect bool) {
	res := solve(a, b, c)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	var x, y, z int
	for i := res[0] - 1; i < res[1]; i++ {
		x += a[i]
	}
	for i := res[2] - 1; i < res[3]; i++ {
		y += b[i]
	}
	for i := res[4] - 1; i < res[5]; i++ {
		z += c[i]
	}
	tot := sum(a)
	tmp := (tot + 2) / 3
	if x < tmp || y < tmp || z < tmp {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 1, 1, 1, 1}
	b := []int{1, 1, 5, 1, 1}
	c := []int{1, 1, 1, 1, 5}
	expect := true
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{5, 6, 1, 2, 3, 4}
	c := []int{3, 4, 5, 6, 1, 2}
	expect := true
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 4, 4, 4}
	b := []int{4, 4, 4, 4}
	c := []int{4, 4, 4, 4}
	expect := false
	runSample(t, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 10, 5, 2, 10}
	b := []int{9, 6, 9, 7, 1}
	c := []int{10, 7, 10, 2, 3}
	expect := false
	runSample(t, a, b, c, expect)
}
