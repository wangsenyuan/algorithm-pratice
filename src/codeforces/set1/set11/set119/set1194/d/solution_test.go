package main

import "testing"

func runSample(t *testing.T, n int, k int, expect string) {
	res := solve(n, k)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 0, 3
	expect := bob
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 3
	expect := alice
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 4
	expect := bob
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 4, 4
	expect := alice
	runSample(t, n, k, expect)
}
