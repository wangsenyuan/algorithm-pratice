package main

import "testing"

func runSample(t *testing.T, P string, friends []string, expect string) {
	res := solve(P, friends)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := "12:01 AM"
	friends := []string{
		"12:00 AM 11:42 PM",
		"12:01 AM 11:59 AM",
		"12:30 AM 12:00 PM",
		"11:59 AM 11:59 PM",
	}
	expect := "1100"
	runSample(t, P, friends, expect)
}

func TestSample2(t *testing.T) {
	P := "04:12 PM"
	friends := []string{
		"12:00 AM 11:59 PM",
		"01:00 PM 04:12 PM",
		"04:12 PM 04:12 PM",
		"04:12 AM 04:12 AM",
		"12:00 PM 11:59 PM",
	}
	expect := "11101"
	runSample(t, P, friends, expect)
}
