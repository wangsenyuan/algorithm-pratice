package main

import "testing"

func runSample(t *testing.T, field []string, expect int) {
	res := int(solve(field))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	field := []string{
		"**",
		"*.",
	}
	expect := 1
	runSample(t, field, expect)
}

func TestSample2(t *testing.T) {
	field := []string{
		"*..*",
		".**.",
		"*.**",
	}
	expect := 9
	runSample(t, field, expect)
}
