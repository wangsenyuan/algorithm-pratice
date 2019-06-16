package p914

import "testing"

func runSample(t *testing.T, deck []int, expect bool) {
	res := hasGroupsSizeX(deck)
	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", deck, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3, 4, 4, 3, 2, 1}, true)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 1, 2, 2, 2, 3, 3}, false)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1}, false)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{1, 1}, true)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{1, 1, 2, 2, 2, 2}, true)
}

func TestSample6(t *testing.T) {
	runSample(t, []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}, true)
}
