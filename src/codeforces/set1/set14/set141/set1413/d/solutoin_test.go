package main

import "testing"

func runSample(t *testing.T, events []string, expect bool) {
	res := solve(events)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
}

func TestSample1(t *testing.T) {
	events := []string{
		"+",
		"+",
		"- 2",
		"+",
		"- 3",
		"+",
		"- 1",
		"- 4",
	}
	expect := true
	runSample(t, events, expect)
}
func TestSample2(t *testing.T) {
	events := []string{
		"- 1",
		"+",
	}
	expect := false
	runSample(t, events, expect)
}

func TestSample3(t *testing.T) {
	events := []string{
		"+",
		"+",
		"+",
		"- 2",
		"- 1",
		"- 3",
	}
	expect := false
	runSample(t, events, expect)
}
