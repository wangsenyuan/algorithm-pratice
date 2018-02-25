package p788

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := rotatedDigits(n)

	if res != expect {
		t.Errorf("sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 12, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 15, 6)
}
