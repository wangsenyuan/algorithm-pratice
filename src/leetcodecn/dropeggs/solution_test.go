package dropeggs

import "testing"

func runSample(t *testing.T, K, N int, expect int) {
	res := superEggDrop(K, N)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", K, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 6, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 14, 4)
}
