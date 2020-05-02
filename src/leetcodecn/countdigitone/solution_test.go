package countdigitone

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := countDigitOne(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 10000, 4001)
}

func TestSample4(t *testing.T) {
	runSample(t, 111, 36)
}
