package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 15
	S := []string{
		"+ 2432543",
		"-",
		"+ 4567886",
		"+ 65638788",
		"-",
		"+ 578943",
		"-",
		"-",
		"+ 62356680",
		"-",
		"+ 711111",
		"-",
		"+ 998244352",
		"-",
		"-",
	}
	expect := 750759115
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := []string{
		"-",
		"+ 1",
		"+ 2",
		"-",
	}
	expect := 16
	runSample(t, n, S, expect)
}
