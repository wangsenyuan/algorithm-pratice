package main

import "testing"

func runSample(t *testing.T, R, C int, G []string, expect string) {
	res := solve(R, C, G)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %s, bug got %s", R, C, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 7, 7
	G := []string{
		"^#^#^^^",
		"^#^^^#^",
		"^#^^^^^",
		"^^^#^^^",
		"^^^^^^^",
		"^^^^^^^",
		"^^^^^^^",
	}

	expect := "Asuna"
	runSample(t, R, C, G, expect)
}
