package main

import (
	"math/rand"
	"sort"
	"testing"
)

func runSample(t *testing.T, p []int) {
	n := len(p)
	q := make([]int, n)
	copy(q, p)
	res := solve(p)

	if len(res) > 5*n {
		t.Fatalf("Sample result %v, took too much times", res)
	}

	n /= 2

	for _, cur := range res {
		i, j := cur[0], cur[1]
		if j-i < n {
			t.Fatalf("Sample result %v, has wrong step %v", res, cur)
		}
		i--
		j--
		q[i], q[j] = q[j], q[i]
	}

	if !sort.IntsAreSorted(q) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 1}
	runSample(t, p)
}

func TestSample2(t *testing.T) {
	p := []int{3, 4, 1, 2}
	runSample(t, p)
}

func TestSample3(t *testing.T) {
	p := []int{2, 5, 3, 1, 4, 6}
	runSample(t, p)
}

func TestSample4(t *testing.T) {
	n := 100
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	q := make([]int, n)
	for k := 0; k < 10; k++ {
		rand.Shuffle(n, func(i, j int) {
			p[i], p[j] = p[j], p[i]
		})
		copy(q, p)
		runSample(t, q)
	}
}
