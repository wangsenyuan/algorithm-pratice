package main

import "testing"

func runSample(t *testing.T, n int, A []byte, m int, B []byte, expect bool) {
	res := solve(n, A, m, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []byte("aB aB Ca Bc")
	m := 6
	B := []byte("aB Ca ca ab ba bc")
	expect := true

	runSample(t, n, A, m, B, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []byte("aB aB Ca Bc")
	m := 3
	B := []byte("Ba aC bc")
	expect := false

	runSample(t, n, A, m, B, expect)
}
