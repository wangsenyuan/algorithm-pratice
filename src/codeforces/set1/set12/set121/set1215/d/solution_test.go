package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0523"
	expect := Bicarp
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "??"
	expect := Bicarp
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "?054??0?"
	expect := Bicarp
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "???00?"
	expect := Monocarp
	runSample(t, s, expect)
}
