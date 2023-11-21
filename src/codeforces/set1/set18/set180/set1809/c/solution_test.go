package main

import "testing"

func runSample(t *testing.T, n int, k int) {
	res := solve(n, k)

	var cnt int
	for i := 0; i < n; i++ {
		var sum int
		for j := i; j < n; j++ {
			sum += res[j]
			if sum > 0 {
				cnt++
			}
		}
	}

	if cnt != k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 30, 100)
}
