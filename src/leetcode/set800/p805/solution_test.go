package p805

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := splitArraySameAverage(A)
	if res != expect {
		t.Errorf("sample %v, expect %t, but got %t", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8}
	runSample(t, A, true)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1}
	runSample(t, A, false)
}

func TestSample3(t *testing.T) {
	A := []int{3, 74, 86, 33, 50, 96, 79, 51, 27, 29, 80, 65, 19, 92, 58, 25, 59, 87, 61, 17, 89, 17}
	runSample(t, A, false)
}

func TestSample4(t *testing.T) {
	A := []int{0, 13, 13, 7, 5, 0, 10, 19, 5}
	runSample(t, A, true)
}

func TestSample5(t *testing.T) {
	A := []int{5, 3, 11, 19, 2}
	runSample(t, A, true)
}

func TestSample6(t *testing.T) {
	A := []int{0, 0, 0}
	runSample(t, A, true)
}

func TestSample7(t *testing.T) {
	A := []int{72, 56, 81, 54, 15, 96, 80, 90, 8}
	runSample(t, A, true)
}
