package main

import "testing"

func runSample(t *testing.T, num string, expect string) {
	res := solve(num)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1"
	expect := "1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "99"
	expect := "100"
	runSample(t, s, expect)
}
