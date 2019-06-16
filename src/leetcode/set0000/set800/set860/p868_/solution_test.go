package p868_

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := binaryGap(N)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 22, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 8, 0)
}
