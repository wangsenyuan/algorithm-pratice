package p1039

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := minScoreTriangulation(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3}, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{3, 7, 4, 5}, 144)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1, 3, 1, 4, 1, 5}, 13)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{2, 1, 4, 4}, 24)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{35, 73, 90, 27, 71, 80, 21, 33, 33, 13, 48, 12, 68, 70, 80, 36, 66, 3, 70, 58}, 140295)
}
