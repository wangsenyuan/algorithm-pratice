package main

import "testing"

func runSample(t *testing.T, records []string, expect int) {
	res := solve(records)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	records := []string{
		"! abc",
		". ad",
		". b",
		"! cd",
		"? c",
	}
	expect := 1
	runSample(t, records, expect)
}

func TestSample2(t *testing.T) {
	records := []string{
		"! hello",
		"! codeforces",
		"? c",
		". o",
		"? d",
		"? h",
		". l",
		"? e",
	}
	expect := 2
	runSample(t, records, expect)
}

func TestSample3(t *testing.T) {
	records := []string{
		"! ababahalamaha",
		"? a",
		"? b",
		"? a",
		"? b",
		"? a",
		"? h",
	}
	expect := 0
	runSample(t, records, expect)
}
