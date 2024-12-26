package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res == expect {
		return
	}

	if expect == "Correct" || expect == "Impossible" || res == "Correct" || res == "Impossible" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "2+2=4"
	expect := "Correct"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "123456789+9876543210=111111110+11-1"
	expect := "123456789+987654321=1111111100+11-1"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10+9=10"
	expect := "Impossible"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "24=55-13"
	expect := "42=55-13"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1000000000-10=9999999999"
	expect := "Impossible"
	runSample(t, s, expect)
}
