package main

import "testing"

func runSample(t *testing.T, a, b string, num string, expect string) {
	res := solve(a, b, num)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "5"
	b := "R"
	num := "4"
	expect := "IV"
	runSample(t, a, b, num, expect)
}

func TestSample2(t *testing.T) {
	a := "10"
	b := "2"
	num := "1"
	expect := "1"
	runSample(t, a, b, num, expect)
}

func TestSample3(t *testing.T) {
	a := "18"
	b := "R"
	num := "GH"
	expect := "CCCV"
	runSample(t, a, b, num, expect)
}

func TestSample4(t *testing.T) {
	a := "23"
	b := "R"
	num := "57F"
	expect := "MMDCCCXXI"
	runSample(t, a, b, num, expect)
}