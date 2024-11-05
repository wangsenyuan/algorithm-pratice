package main

import "testing"

func runSample(t *testing.T, cake []string, expect int) {
	res := solve(cake)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	cake := []string{
		"*..",
		".*.",
		"..*",
		"...",
	}
	expect := 3
	runSample(t, cake, expect)
}

func TestSample2(t *testing.T) {
	cake := []string{
		".*..",
		"*...",
		"...*",
		"..*.",
	}
	expect := 2
	runSample(t, cake, expect)
}

func TestSample3(t *testing.T) {
	cake := []string{
		"..**",
		"*...",
		"....",
	}
	expect := 1
	runSample(t, cake, expect)
}

func TestSample4(t *testing.T) {
	cake := []string{
		"..*..",
		".....",
		"**...",
		"**...",
		"**...",
	}
	expect := 1
	runSample(t, cake, expect)
}
