package main

import "testing"

func runSample(t *testing.T, room []string, expect int) {
	res := solve(room)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	room := []string{
		"000",
		"010",
		"000",
	}
	expect := 8
	runSample(t, room, expect)
}

func TestSample2(t *testing.T) {
	room := []string{
		"1100",
		"0000",
		"0000",
		"0000",
		"0000",
	}
	expect := 16
	runSample(t, room, expect)
}
