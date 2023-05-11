package main

import "testing"

func runSample(t *testing.T, fields []string, cnt int, expect int) {
	cnt1, res := solve(fields)

	if cnt != cnt1 || expect != res {
		t.Fatalf("Sample expect %d %d, but got %d %d", cnt, expect, cnt1, res)
	}
}

func TestSample1(t *testing.T) {
	field := []string{
		"...",
		".*.",
	}
	cnt := 3
	expect := 1

	runSample(t, field, cnt, expect)
}

func TestSample2(t *testing.T) {
	field := []string{
		"...",
		".*.",
		"...",
	}
	cnt := 4
	expect := 2

	runSample(t, field, cnt, expect)
}
