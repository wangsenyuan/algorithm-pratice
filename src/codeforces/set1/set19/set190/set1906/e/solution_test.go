package main

import "testing"

func runSample(t *testing.T, c []int, expect bool) {
	ok, a, b := solve(c)
	if expect != ok {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}
	for i, j, k := 0, 0, 0; k < len(c); k++ {
		var d int
		if j == len(b) || i < len(a) && a[i] < b[j] {
			d = a[i]
			i++
		} else {
			d = b[j]
			j++
		}
		if d != c[k] {
			t.Fatalf("Sample result %v %v, not correct", a, b)
		}
	}
}

func TestSample1(t *testing.T) {
	c := []int{3, 1, 4, 5, 2, 6}
	expect := true
	runSample(t, c, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expect := true
	runSample(t, c, expect)
}

func TestSample3(t *testing.T) {
	c := []int{4, 3, 2, 1}
	expect := false
	runSample(t, c, expect)
}
