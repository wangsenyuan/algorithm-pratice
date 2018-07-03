package main

import "testing"

func runSample(t *testing.T, S []string, Q []string, expect int) {
	res := solve(S, Q)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", S, Q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []string{"Yeehaw",
		"NSM",
		"Dont Ask",
		"B9",
		"Googol",
	}
	Q := []string{
		"Yeehaw",
		"Yeehaw",
		"Googol",
		"B9",
		"Googol",
		"NSM",
		"B9",
		"NSM",
		"Dont Ask",
		"Googol",
	}
	expect := 1
	runSample(t, S, Q, expect)
}

func TestSample2(t *testing.T) {
	S := []string{"Yeehaw",
		"NSM",
		"Dont Ask",
		"B9",
		"Googol",
	}
	Q := []string{
		"Googol",
		"Dont Ask",
		"NSM",
		"NSM",
		"Yeehaw",
		"Yeehaw",
		"Googol",
	}
	expect := 0
	runSample(t, S, Q, expect)
}
