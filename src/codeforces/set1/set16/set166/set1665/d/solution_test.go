package main

import (
	rand2 "math/rand"
	"testing"
)

func runSample(t *testing.T, x int) {
	ask := func(a, b int) int {
		return gcd(a+x, b+x)
	}

	ans := solve(ask)

	if ans != x {
		t.Fatalf("Sample expect %d, but got %d", x, ans)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000000000)
}

func TestSample3(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x := (rand2.Int31() >> 2) + 1
		runSample(t, int(x))
	}
}
