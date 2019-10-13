package p1223

import "testing"

func runSample(t *testing.T, n int, rollMax []int, expect int) {
	res := dieSimulator(n, rollMax)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, rollMax, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []int{1, 1, 2, 2, 2, 3}, 34)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, []int{1, 1, 1, 1, 1, 1}, 30)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{1, 1, 1, 2, 2, 3}, 181)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, []int{2, 1, 1, 3, 3, 2}, 1082)
}
