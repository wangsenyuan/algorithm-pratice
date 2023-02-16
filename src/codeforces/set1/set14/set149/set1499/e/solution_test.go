package main

import (
	"testing"
)

func runSample(t *testing.T, x, y string, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := "aaa"
	y := "bb"
	expect := 24
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := "code"
	y := "forces"
	expect := 1574
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := "aaaaa"
	y := "aaa"
	expect := 0
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := "justamassivetesttocheck"
	y := "howwellyouhandlemodulooperations"
	expect := 667387032
	runSample(t, x, y, expect)
}
