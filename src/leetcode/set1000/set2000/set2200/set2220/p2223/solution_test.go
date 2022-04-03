package p2223

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(sumScores(s))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "babab"
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "azbazbzaz"
	expect := 14
	runSample(t, s, expect)
}