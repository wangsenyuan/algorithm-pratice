package main

import "testing"

func runSample(t *testing.T, a [][]int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if res[i][j] != a[i][j] && res[i][j] != a[i][j]+1 {
				t.Fatalf("Sample result %v, having wrong elment at %d %d", res, i, j)
			}
			if i > 0 && res[i][j] == res[i-1][j] {
				t.Fatalf("Sample result %v, having wrong elment at %d %d", res, i, j)
			}
			if j > 0 && res[i][j] == res[i][j-1] {
				t.Fatalf("Sample result %v, having wrong elment at %d %d", res, i, j)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	expect := true

	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]int{
		{1, 1},
		{3, 3},
	}
	expect := true

	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := [][]int{
		{1, 3},
		{2, 2},
	}
	expect := true

	runSample(t, a, expect)
}
