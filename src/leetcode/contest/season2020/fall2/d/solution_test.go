package d

import "testing"

func runSample(t *testing.T, k, n int, expect int) {
	res := keyboard(k, n)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", k, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 26)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 650)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 100, 629445915)
}
