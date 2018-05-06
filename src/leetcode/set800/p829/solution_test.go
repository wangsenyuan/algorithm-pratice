package p829

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := consecutiveNumbersSum(N)

	if res != expect {
		t.Errorf("sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 15, 4)
}
