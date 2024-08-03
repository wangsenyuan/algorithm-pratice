package main

import "testing"

func runSample(t *testing.T, days []string, expect string) {
	res := solve(days)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	days := []string{
		"win 10",
		"win 5",
		"win 3",
		"sell 5",
		"sell 3",
		"win 10",
		"sell 10",
	}
	expect := "1056"
	runSample(t, days, expect)
}
