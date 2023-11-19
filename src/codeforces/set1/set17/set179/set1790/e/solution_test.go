package main

import "testing"

func runSample(t *testing.T, x int, expect bool) {
	res := solve(x)

	if (len(res) > 0) != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	a, b := res[0], res[1]
	if a^b != x || (a+b)/2 != x {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	x := 2
	expect := true
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := 5
	expect := false
	runSample(t, x, expect)
}

func TestSample3(t *testing.T) {
	x := 10
	expect := true
	runSample(t, x, expect)
}

func TestSample4(t *testing.T) {
	x := 6
	expect := false
	runSample(t, x, expect)
}

func TestSample5(t *testing.T) {
	x := 18
	expect := true
	runSample(t, x, expect)
}

func TestSample6(t *testing.T) {
	x := 36
	expect := true
	runSample(t, x, expect)
}
