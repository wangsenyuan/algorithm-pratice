package p892

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := surfaceArea(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{{2}}
	expect := 10
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	expect := 34
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{{1, 0}, {0, 2}}
	expect := 16
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	expect := 32
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	expect := 46
	runSample(t, A, expect)
}
