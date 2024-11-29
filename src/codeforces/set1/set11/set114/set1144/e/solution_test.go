package main

import "testing"

func runSample(t *testing.T, s string, x string, expect string) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "az"
	x := "bf"
	expect := "bc"
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "afogk"
	x := "asdji"
	expect := "alvuw"
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "nijfvj"
	x := "tvqhwp"
	expect := "qoztvz"
	runSample(t, s, x, expect)
}
