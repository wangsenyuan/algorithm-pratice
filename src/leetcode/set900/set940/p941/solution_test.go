package p941

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := validMountainArray(A)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 1}, false)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{3, 5, 5}, false)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 3, 2, 1}, true)
}
func TestSample4(t *testing.T) {
	runSample(t, []int{3, 2, 1}, false)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, false)
}
