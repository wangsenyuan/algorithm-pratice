package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect string) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 15, 4, "Timur")
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 9, 5, "Marsel")
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 2, 1, "Timur")
}

func TestSample4(t *testing.T) {
	runSample(t, 225, 187, 20, "Marsel")
}
