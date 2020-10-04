package p1546

import "testing"

func runSample(t *testing.T, n int, cuts []int, expect int) {
	res := minCost(n, cuts)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, cuts, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	cuts := []int{1, 3, 4, 5}
	expect := 16
	runSample(t, n, cuts, expect)
}

func TestSample2(t *testing.T) {
	n := 36
	cuts := []int{13, 17, 15, 18, 3, 22, 27, 6, 35, 7, 11, 28, 26, 20, 4, 5, 21, 10, 8}
	expect := 150
	runSample(t, n, cuts, expect)
}
