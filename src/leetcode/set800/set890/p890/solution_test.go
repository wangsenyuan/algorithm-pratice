package p890

import "testing"

func runSample(t *testing.T, N int, dislikes [][]int, expect bool) {
	res := possibleBipartition(N, dislikes)
	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", N, dislikes, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 4
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	expect := true
	runSample(t, N, dislikes, expect)
}

func TestSample2(t *testing.T) {
	N := 3
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 3}}
	expect := false
	runSample(t, N, dislikes, expect)
}

func TestSample3(t *testing.T) {
	N := 5
	dislikes := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	expect := false
	runSample(t, N, dislikes, expect)
}
