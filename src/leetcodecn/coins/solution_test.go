package coins

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := waysToChange(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000, 142511)
}
