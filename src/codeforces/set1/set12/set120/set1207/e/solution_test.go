package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, x int) {
	ask := func(arr []int) int {
		n := len(arr)
		i := rand.Intn(n)
		return arr[i] ^ x
	}

	res := solve(ask)

	if res != x {
		t.Fatalf("Sample getting wrong result %d, instead of %d", res, x)
	}
}

func TestSample(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := rand.Intn(1 << 14)
		runSample(t, x)
	}
}
