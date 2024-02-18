package p3040

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := maxSelectedElements(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 5, 1, 1}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 4, 7, 10}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{8, 10, 6, 12, 9, 12, 2, 3, 13, 19, 11, 18, 10, 16}
	expect := 8
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{8, 13, 18, 10, 16, 19, 11, 17, 15, 18, 9, 12, 15, 8, 9, 14, 7}
	expect := 14
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{12, 11, 8, 7, 2, 10, 18, 12}
	expect := 6
	runSample(t, a, expect)
}
