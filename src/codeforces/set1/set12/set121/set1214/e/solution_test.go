package main

import "testing"

func runSample(t *testing.T, d []int) {
	res := solve(d)

	n := len(d)
	g := NewGraph(2*n, 4*n)

	for _, e := range res {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := g.prepare()

	for i := 0; i < n; i++ {
		u := 2 * i
		v := u + 1
		di := dist(u, v)
		if di != d[i] {
			t.Fatalf("Sample result %v, not correct, expect %d, but got %d", res, d[i], di)
		}
	}
}

func TestSample1(t *testing.T) {
	d := []int{2, 2, 2}
	runSample(t, d)
}

func TestSample2(t *testing.T) {
	d := []int{2, 2, 2, 1}
	runSample(t, d)
}

func TestSample3(t *testing.T) {
	d := []int{2, 2, 2, 2, 2, 2}
	runSample(t, d)
}

func TestSample4(t *testing.T) {
	d := []int{1, 1}
	runSample(t, d)
}

func TestSample5(t *testing.T) {
	d := []int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	runSample(t, d)
}
