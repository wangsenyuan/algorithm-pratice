package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for i, cur := range res {
		u, v := cur[0], cur[1]
		u--
		v--
		if abs(a[u]-a[v])%(i+1) != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 2, 31, 44, 73}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{52, 63, 25, 21, 5}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{33, 40, 3, 11, 31, 43, 37, 8, 50, 5, 12, 22}
	expect := true
	runSample(t, a, expect)
}
