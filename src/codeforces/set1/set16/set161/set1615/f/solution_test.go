package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "00"
	B := "11"
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "???"
	B := "???"
	expect := 16
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "??1"
	B := "0?0"
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := "??0?"
	B := "??11"
	expect := 14
	runSample(t, A, B, expect)
}

func TestSample5(t *testing.T) {
	A := "?????"
	B := "0??1?"
	expect := 101
	runSample(t, A, B, expect)
}

func TestSample6(t *testing.T) {
	A := "?01??01?1?"
	B := "??100?1???"
	expect := 1674
	runSample(t, A, B, expect)
}
