package p0002

import "testing"

func runSample(t *testing.T, m, n int, expect int) {
	res := rangeBitwiseAnd(m, n)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", m, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 7, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 6, 0)
}

func TestSample4(t *testing.T) {
	for x := 1; x < 100; x++ {
		var expect = x
		for y := x; y < 1000; y++ {
			expect &= y
			runSample(t, x, y, expect)
		}
	}
}
