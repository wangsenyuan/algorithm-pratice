package main

import "testing"

func runSample(t *testing.T, x string, y string, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := "abcde"
	y := "abxde"
	expect := 1
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := "hello"
	y := "olleo"
	expect := 2
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := "ab"
	y := "cd"
	expect := 3
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := "aaaaaaa"
	y := "abbbbba"
	expect := 9
	runSample(t, x, y, expect)
}

func TestSample5(t *testing.T) {
	x := "q"
	y := "q"
	expect := 0
	runSample(t, x, y, expect)
}

func TestSample6(t *testing.T) {
	x := "yoyoyo"
	y := "oyoyoy"
	expect := 2
	runSample(t, x, y, expect)
}

func TestSample7(t *testing.T) {
	x := "abcdefgh"
	y := "hguedfbh"
	expect := 6
	runSample(t, x, y, expect)
}
