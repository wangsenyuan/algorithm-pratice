package main

import "testing"

func runSample(t *testing.T, land []string, expect int) {
	res := solve(land)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	land := []string{
		"GWGGW",
		"GGWGG",
		"GWGGG",
		"WGGGG",
	}
	expect := 11
	runSample(t, land, expect)
}

func TestSample2(t *testing.T) {
	land := []string{
		"GWW",
		"WWW",
		"WWG",
	}
	expect := 7
	runSample(t, land, expect)
}

func TestSample3(t *testing.T) {
	land := []string{
		"G",
	}
	expect := 0
	runSample(t, land, expect)
}
