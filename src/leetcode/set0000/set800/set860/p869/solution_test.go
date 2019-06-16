package p869

import "testing"

func runSample(t *testing.T, N int, expect bool) {
	res := reorderedPowerOf2(N)

	if res != expect {
		t.Errorf("sample %d, expect %t, but got %t", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 16, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 24, false)
}

func TestSample5(t *testing.T) {
	runSample(t, 46, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 5792021, true)
}
