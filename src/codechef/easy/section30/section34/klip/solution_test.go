package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	k := 2
	expect := "000"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1011"
	k := 2
	expect := "0001"
	runSample(t, s, k, expect)
}
func TestSample3(t *testing.T) {
	s := "01010"
	k := 3
	expect := "00011"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "01010"
	k := 2
	expect := "00000"
	runSample(t, s, k, expect)
}