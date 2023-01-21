package main

import (
	"testing"
)

func runSample(t *testing.T, n int, x int) {
	var cnt int
	ask := func(c int) int {
		cnt++
		if cnt > 10 {
			panic("more than 10 times query")
		}
		x += c
		return x / n
	}

	res := solve(n, ask)

	if res != x {
		t.Errorf("Sample expect %d, but got %d", x, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	x := 2
	runSample(t, n, x)
}

func TestSample2(t *testing.T) {
	n := 5
	x := 2
	runSample(t, n, x)
}


func TestSample3(t *testing.T) {
	n := 1000
	x := 10
	runSample(t, n, x)
}

func TestSample4(t *testing.T) {
	n := 10
	x := 2
	runSample(t, n, x)
}
