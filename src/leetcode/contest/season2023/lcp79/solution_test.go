package lcp79

import "testing"

func runSample(t *testing.T, mat []string, mantra string, expect int) {
	res := extractMantra(mat, mantra)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"sd",
		"ep",
	}
	mantra := "speed"
	expect := 10
	runSample(t, mat, mantra, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"abc", "daf", "geg",
	}
	mantra := "sad"
	expect := -1
	runSample(t, mat, mantra, expect)
}
