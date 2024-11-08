package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	_, lpf := getPrimes(n * 2)

	m := len(res)
	if lpf[m] != m {
		t.Fatalf("Sample result %v, not correct, as its count is not a prime", res)
	}

	deg := make([]int, n+1)

	for _, e := range res {
		u, v := e[0], e[1]
		deg[u]++
		deg[v]++
	}

	for i := 1; i <= n; i++ {
		if lpf[deg[i]] != deg[i] {
			t.Fatalf("Sample result %v, not correct, as deg of %d is %d, not a prime", res, i, deg[i])
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 12)
}


func TestSample3(t *testing.T) {
	runSample(t, 108)
}
