package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "BRBBW"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "B"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "WB"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "BRB"
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "RBB"
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "WWWWWWW"
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "RBWBWRRBW"
	expect := false
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "BRBRBRBRRB"
	expect := true
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "BBBRWWRRRWBR"
	expect := false
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := "BRBRBRBRBW"
	expect := true
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := "RBRBWBBBRRWWBR"
	expect := true
	runSample(t, s, expect)
}
