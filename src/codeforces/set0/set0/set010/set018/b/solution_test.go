package main

import "testing"

func runSample(t *testing.T, n int, d int, m int, l int, expect int) {
	res := solve(n, d, m, l)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d, m, l := 2, 2, 5, 3
	expect := 4
	runSample(t, n, d, m, l, expect)
}

func TestSample2(t *testing.T) {
	n, d, m, l := 5, 4, 11, 8
	expect := 20
	runSample(t, n, d, m, l, expect)
}

func TestSample3(t *testing.T) {
	n, d, m, l := 228385, 744978, 699604, 157872
	expect := 2979912
	runSample(t, n, d, m, l, expect)
}

func TestSample4(t *testing.T) {
	n, d, m, l := 3, 6, 6, 3
	expect := 18
	runSample(t, n, d, m, l, expect)
}

func TestSample5(t *testing.T) {
	n, d, m, l := 1000000, 1, 1000000, 999999
	expect := 1000000000000
	runSample(t, n, d, m, l, expect)
}
