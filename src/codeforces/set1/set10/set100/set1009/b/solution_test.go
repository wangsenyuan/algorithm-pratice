package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "100210"
	expect := "001120"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "11222121"
	expect := "11112222"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "20"
	expect := "20"
	runSample(t, s, expect)
}
