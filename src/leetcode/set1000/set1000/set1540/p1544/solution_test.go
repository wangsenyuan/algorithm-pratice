package p1544

import "testing"

func runSample(t *testing.T, n int, k int, expect byte) {
	res := findKthBit(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %c, but got %c", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, '0')
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 11, '1')
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, '1')
}
