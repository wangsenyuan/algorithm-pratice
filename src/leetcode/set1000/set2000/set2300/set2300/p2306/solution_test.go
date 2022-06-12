package p2306

import "testing"

func runSample(t *testing.T, ideas []string, expect int64) {
	res := distinctNames(ideas)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ideas := []string{
		"coffee", "donuts", "time", "toffee",
	}
	var expect int64 = 6
	runSample(t, ideas, expect)
}

func TestSample2(t *testing.T) {
	ideas := []string{
		"lack", "back",
	}
	var expect int64 = 0
	runSample(t, ideas, expect)
}
