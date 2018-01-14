package p764

import "testing"

func runSample(t *testing.T, N int, mines [][]int, expect int) {
	res := orderOfLargestPlusSign(N, mines)
	if res != expect {
		t.Errorf("sample %d %v, should give %d, but got %d", N, mines, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	mines := [][]int{
		{4, 2},
	}
	expect := 2
	runSample(t, N, mines, expect)
}

func TestSample2(t *testing.T) {
	N := 2
	mines := [][]int{}
	expect := 1
	runSample(t, N, mines, expect)
}

func TestSample3(t *testing.T) {
	N := 1
	mines := [][]int{{0, 0}}
	expect := 0
	runSample(t, N, mines, expect)
}

func TestSample4(t *testing.T) {
	N := 3
	mines := [][]int{{0, 0}}
	expect := 2
	runSample(t, N, mines, expect)
}
