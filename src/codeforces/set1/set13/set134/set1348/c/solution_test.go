package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "baba"
	k := 2
	expect := "ab"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "baacb"
	k := 2
	expect := "abbc"
	// abbc
	// a
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "baacb"
	k := 3
	expect := "b"
	// abbc
	// a
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "baacb"
	k := 4
	expect := "b"
	// abbc
	// a
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "baacb"
	k := 5
	expect := "c"
	// abbc
	// a
	runSample(t, s, k, expect)
}
