package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, a []int, expect int64) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample %v expect %d, but got %d", a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 5}
	var expect int64 = 14
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(1 << 20)
	}
	var expect int64 = bruteForce(a)
	runSample(t, a, expect)
}
