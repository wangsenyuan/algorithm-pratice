package main

import "testing"

func runSample(t *testing.T, a []string, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []string{
		"nn",
		"aa",
		"rr",
		"ee",
		"kk",
	}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []string{
		"narek",
	}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []string{
		"nare",
	}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []string{
		"nrrarek",
		"nrnekan",
		"uuuuuuu",
		"ppppppp",
		"nkarekz",
	}
	expect := 7
	runSample(t, a, expect)
}
