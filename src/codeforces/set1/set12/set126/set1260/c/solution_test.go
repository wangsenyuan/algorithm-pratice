package main

import "testing"

func runSample(t *testing.T, r int, b int, k int, expect string) {
	res := solve(r, b, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 2, obey)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 10, 4, rebel)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 2, 3, obey)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 2, 2, obey)
}
