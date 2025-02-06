package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int) {
	n := len(p)

	var cnt int

	ask := func(w int, i int, j int, x int) int {
		cnt++
		if cnt > 3*n/2+30 {
			t.Fatalf("Sample solution asked too much")
		}
		i--
		j--
		if w == 1 {
			return max(min(x, p[i]), min(x+1, p[j]))
		}
		return min(max(x, p[i]), max(x+1, p[j]))
	}

	res := solve(n, ask)

	if !reflect.DeepEqual(res, p) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{3, 1, 4, 2}
	runSample(t, p)
}

func TestSample2(t *testing.T) {
	n := 1001
	p := make([]int, n)
	for i := range n {
		p[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
	runSample(t, p)
}

func TestSample3(t *testing.T) {
	n := 10000
	p := make([]int, n)
	for i := range n {
		p[i] = i + 1
	}
	rand.Shuffle(n, func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
	runSample(t, p)
}
