package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "abb"
	B := "baz"
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "zzc"
	B := "aaa"
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "fxbs"
	B := "dkrc"
	expect := 11
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := "eaufq"
	B := "drtkn"
	expect := 9
	runSample(t, A, B, expect)
}
