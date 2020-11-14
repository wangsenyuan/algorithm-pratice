package main

import "testing"

func runSample(t *testing.T, n int, s string) {
	w := solve(n, s)

	for i := 0; i < n; i++ {
		if w[i] != '0' && w[i] != '1' {
			t.Fatalf("Sample result %s, not correct", w)
		}
		if !similar(w, s[i:i+n]) {
			t.Fatalf("Sample %s, result %s, not correct with %s", s, w, s[i:i+n])
		}
	}
}

func similar(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			return true
		}
	}
	return false
}

func TestSample1(t *testing.T) {
	n := 1
	s := "1"
	runSample(t, n, s)
}

func TestSample2(t *testing.T) {
	n := 3
	s := "00000"
	runSample(t, n, s)
}

func TestSample3(t *testing.T) {
	n := 4
	s := "1110000"
	runSample(t, n, s)
}

func TestSample4(t *testing.T) {
	n := 2
	s := "101"
	runSample(t, n, s)
}
