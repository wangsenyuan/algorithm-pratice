package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect bool) {
	res := solve(a, b)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != res[i]|res[i+1] || b[i] != res[i]&res[i+1] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 3, 2}
	b := []int{1, 2, 0}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3}
	b := []int{3, 2}
	expect := false
	runSample(t, a, b, expect)
}
