package p1703

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := halvesAreAlike(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "book", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "textbook", false)
}
func TestSample3(t *testing.T) {
	runSample(t, "MerryChristmas", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "AbCdEfGh", true)
}