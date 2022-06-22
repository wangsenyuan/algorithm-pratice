package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, P []int) {
	n := len(P)
	p := make([]int, n)
	copy(p, P)
	res := solve(P)
	if len(res) > n*n {
		t.Fatalf("Sample result %v, TTL", res)
	}

	x := f(p)

	for _, op := range res {
		i, j := op[0], op[1]
		i--
		j--
		p[i], p[j] = p[j], p[i]
		y := f(p)
		if y > x {
			t.Fatalf("Sample result %v, after swap %v, get more value %d than %d", res, op, y, x)
		}
	}

	if !sort.IntsAreSorted(p) {
		t.Errorf("Sample result %v, get not sorted result %v", res, p)
	}
}

func f(p []int) int {
	var res int
	n := len(p)
	for i := 0; i < n-1; i++ {
		if p[i] > p[i+1] {
			res++
		}
	}
	res++
	return res
}

func TestSample1(t *testing.T) {
	P := []int{3, 2, 1}
	runSample(t, P)
}

func TestSample2(t *testing.T) {
	P := []int{1, 2}
	runSample(t, P)
}

func TestSample3(t *testing.T) {
	P := []int{4, 1, 2, 3}
	runSample(t, P)
}
