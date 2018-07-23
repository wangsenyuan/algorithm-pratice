package p838

import "testing"

func runSample(t *testing.T, dominoes string, expect string) {
	res := pushDominoes(dominoes)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", dominoes, expect, res)
	}
}

func TestSample(t *testing.T) {
	dominoes := ".L.R...LR..L.."
	expect := "LL.RR.LLRRLL.."
	runSample(t, dominoes, expect)
}

func TestSample2(t *testing.T) {
	dominoes := "RR.L"
	expect := "RR.L"
	runSample(t, dominoes, expect)
}
