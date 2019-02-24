package p997

import "testing"

func runSample(t *testing.T, N int, trust [][]int, expect int) {
	res := findJudge(N, trust)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, trust, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 2
	trust := [][]int{{1, 2}}
	expect := 2
	runSample(t, N, trust, expect)
}

func TestSample2(t *testing.T) {
	N := 3
	trust := [][]int{{1, 3}, {2, 3}}
	expect := 3
	runSample(t, N, trust, expect)
}

func TestSample3(t *testing.T) {
	N := 3
	trust := [][]int{{1, 2}, {2, 3}}
	expect := -1
	runSample(t, N, trust, expect)
}

func TestSample4(t *testing.T) {
	N := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	expect := 3
	runSample(t, N, trust, expect)
}
