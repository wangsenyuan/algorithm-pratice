package main

import "testing"

func runSample(t *testing.T, a int, b int, expect string) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 1, 10
	expect := CHEFINA
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 4, 2
	expect := CHEF
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 2, 2
	expect := CHEFINA
	runSample(t, a, b, expect)
}
