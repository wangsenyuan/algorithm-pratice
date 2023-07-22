package main

import "testing"

func runSample(t *testing.T, s string, k int, expect bool) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 4
	s := "100110"
	expect := true
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	s := "1?1"
	expect := true
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	s := "1?0"
	expect := false
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	k := 4
	s := "????"
	expect := true
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	k := 4
	s := "1?0??1?"
	expect := true
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	k := 10
	s := "11??11??11"
	expect := false
	runSample(t, s, k, expect)
}

func TestSample7(t *testing.T) {
	k := 2
	s := "1??1"
	expect := false
	runSample(t, s, k, expect)
}

func TestSample8(t *testing.T) {
	k := 2
	s := "????00"
	expect := false
	runSample(t, s, k, expect)
}
