package p779

import "testing"

func runSample(t *testing.T, N int, K int, expect int) {
	res := kthGrammar(N, K)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", N, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 5, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 30, 434991989, 0)
}

func TestSample6(t *testing.T) {
	runSample(t, 3, 3, 1)
}
