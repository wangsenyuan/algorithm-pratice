package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "04829"
	expect := "02599"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01"
	expect := "01"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "314752277691991"
	expect := "111334567888999"
	runSample(t, s, expect)
}
