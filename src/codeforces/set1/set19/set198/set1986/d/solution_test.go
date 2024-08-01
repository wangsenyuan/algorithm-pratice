package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10"
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "00"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "901"
	expect := 9
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "23311"
	expect := 19
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "987009"
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "99999999999999999999"
	expect := 261
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "1112"
	expect := 12
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "19811678487321784121"
	expect := 93
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "1121"
	expect := 12
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := "2221"
	expect := 24
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := "101"
	expect := 1
	runSample(t, s, expect)
}
