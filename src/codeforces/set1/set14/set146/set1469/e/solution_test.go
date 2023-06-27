package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0110"
	k := 2
	expect := "11"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1001"
	k := 2
	expect := "00"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "0101110001"
	k := 3
	expect := ""
	runSample(t, s, k, expect)
}
