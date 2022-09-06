package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "BRB"
	expect := BOB
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "RRBBB"
	expect := BOB
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "RBRBRB"
	expect := ALICE
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "BBRRBRRB"
	expect := ALICE
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "BRRBRB"
	expect := ALICE
	runSample(t, s, expect)
}
func TestSample6(t *testing.T) {
	s := "RBRBRBRBRRBB"
	expect := ALICE
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "RBRBRBRBBBRR"
	expect := BOB
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "RBBR"
	expect := BOB
	runSample(t, s, expect)
}
