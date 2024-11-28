package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	for i, j := range res {
		if i > 0 && a[j-1] != a[res[i-1]-1]+1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 3, 4, 7, 5, 6, 8}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 5, 2, 4, 6}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 9, 8, 7}
	expect := 1
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{6, 7, 8, 3, 4, 5, 9, 10, 11}
	expect := 6
	runSample(t, a, expect)
}
