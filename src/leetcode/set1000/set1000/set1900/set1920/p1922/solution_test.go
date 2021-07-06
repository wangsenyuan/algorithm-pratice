package p1922

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := countGoodNumbers(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 50, 564908303)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 400)
}

func TestSample3(t *testing.T) {
	runSample(t, 51, 824541501)
}
