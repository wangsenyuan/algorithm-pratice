package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1,2 ,3,...,     10"
	expect := "1, 2, 3, ..., 10"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1,,,4...5......6"
	expect := "1, , , 4 ...5 ... ...6"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "...,1,2,3,..."
	expect := "..., 1, 2, 3, ..."
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1 4 5    6 7 999 1 1    1 2 311111111111111111111111111111111111111111"
	expect := "1 4 5 6 7 999 1 1 1 2 311111111111111111111111111111111111111111"
	runSample(t, s, expect)
}
