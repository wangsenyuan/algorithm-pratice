package main

import "testing"

func runSample(t *testing.T, n int, k uint64, expect uint64) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 64, 11047805202224836936, 1337369305470044825)
}
