package main

import "testing"

func runSample(t *testing.T, cards []string, k int, expect int) {
	res := solve(cards, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cards := []string{
		"SET",
		"ETS",
		"TSE",
	}
	k := 3
	expect := 1
	runSample(t, cards, k, expect)
}

func TestSample2(t *testing.T) {
	cards := []string{
		"SETE",
		"ETSE",
		"TSES",
	}
	k := 4
	expect := 0
	runSample(t, cards, k, expect)
}

func TestSample3(t *testing.T) {
	cards := []string{
		"SETT",
		"TEST",
		"EEET",
		"ESTE",
		"STES",
	}
	k := 4
	expect := 2
	runSample(t, cards, k, expect)
}
