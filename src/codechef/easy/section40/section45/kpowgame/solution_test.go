package main

import "testing"

func runSample(t *testing.T, n int, k int, expect string) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	expect := Shivansh
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	expect := Tina
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 0, 10000
	expect := Tina
	runSample(t, n, k, expect)
}
