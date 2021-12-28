package main

import "testing"

func runSample(t *testing.T, k int, X string, expect string) {
	res := solve(len(X), k, X)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := "faberty"
	k := 3
	expect := "abferty"
	runSample(t, k, X, expect)
}
func TestSample2(t *testing.T) {
	X := "bbaaaba"
	k := 5
	expect := "aababba"
	runSample(t, k, X, expect)
}
func TestSample3(t *testing.T) {
	X := "zxwy"
	k := 4
	expect := "wxyz"
	runSample(t, k, X, expect)
}