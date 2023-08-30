package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(solve(s))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "20230322"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0112223333444445555556666666777777778888888889999999999"
	expect := 185
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "3141592653589793238462643383279502884197169399375105820974944"
	expect := 9
	runSample(t, s, expect)
}
