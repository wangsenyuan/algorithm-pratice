package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	S := []string{
		"*",
		".",
	}
	expect := 0
	runSample(t, n, S, expect)
}
func TestSample2(t *testing.T) {
	n := 2
	S := []string{
		".*",
		"**",
	}
	expect := 2
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	S := []string{
		"*.*",
		".*.",
	}
	expect := 3
	runSample(t, n, S, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	S := []string{
		"**...",
		"...**",
	}
	expect := 5
	runSample(t, n, S, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	S := []string{
		"**.*",
		"**..",
	}
	expect := 5
	runSample(t, n, S, expect)
}

/*
*.
..
*/

func TestSample6(t *testing.T) {
	n := 2
	S := []string{
		"*.",
		"..",
	}
	expect := 0
	runSample(t, n, S, expect)
}

func TestSample7(t *testing.T) {
	n := 2
	S := []string{
		".*",
		"..",
	}
	expect := 0
	runSample(t, n, S, expect)
}
