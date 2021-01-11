package p1718

import "testing"

func runSample(t *testing.T, pairs [][]int, expect int) {
	res := checkWays(pairs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{{1, 2}, {2, 3}}
	expect := 1
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{{1, 2}, {2, 3}, {1, 3}}
	expect := 2
	runSample(t, pairs, expect)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}
	expect := 0
	runSample(t, pairs, expect)
}

func TestSample4(t *testing.T) {
	pairs := [][]int{{4, 5}, {3, 4}, {2, 4}}
	expect := 1
	runSample(t, pairs, expect)
}

func TestSample5(t *testing.T) {
	pairs := [][]int{{3, 4}, {2, 3}, {4, 5}, {2, 4}, {2, 5}, {1, 5}, {1, 4}}
	expect := 0
	runSample(t, pairs, expect)
}
