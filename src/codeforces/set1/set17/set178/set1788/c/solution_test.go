package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	vis := make([]bool, 2*n+1)

	for _, x := range res {
		vis[x[0]] = true
		vis[x[1]] = true

	}

	for i := 1; i <= 2*n; i++ {
		if !vis[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	n := 3
	expect := true
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	expect := true
	runSample(t, n, expect)
}
