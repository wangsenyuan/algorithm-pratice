package p1626

import "testing"

func runSample(t *testing.T, s string, a int, b int, expect string) {
	res := findLexSmallestString(s, a, b)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "5525", 9, 2, "2050")
}

func TestSample2(t *testing.T) {
	runSample(t, "43987654", 7, 3, "00553311")
}
