package main

import (
	"testing"
)

func runSample(t *testing.T, x int) {
	var cnt int
	ask := func(a int, b int) int {
		cnt++
		if cnt > 7 {
			t.Fatalf("Sample asked too much")
		}
		if a >= x {
			a++
		}
		if b >= x {
			b++
		}
		return a * b
	}

	res := solve(ask)

	if res != x {
		t.Fatalf("Sample expect %d, but got %d", x, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	for i := 2; i <= 999; i++ {
		runSample(t, i)
	}
}

func TestSample3(t *testing.T) {
	runSample(t, 25)
}
