package main

import "testing"

func runSample(t *testing.T, n, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 3
	s := "chehefc"
	expect := 0
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 4
	s := "chehefch"
	expect := -1
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 19, 24
	s := "ccccchhhhheeeeeffff"
	expect := 9
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 20, 5
	s := "echhhfhfcecfhechfcch"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample5(t *testing.T) {
	n, k := 99, 27
	s := "ffhffhffcfechchccccfhhhfhhhhhhehhhehhhhheeeeefeeeeeeheeheeeeffcfffffffffffhffffchffeffcefefhhfcehfe"
	expect := 9
	runSample(t, n, k, s, expect)
}
