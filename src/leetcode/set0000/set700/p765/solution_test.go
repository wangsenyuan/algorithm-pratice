package p765

import "testing"

func runSample(t *testing.T, row []int, expect int) {
	cp := make([]int, len(row))
	copy(cp, row)
	res := minSwapsCouples(cp)
	if expect != res {
		t.Errorf("sample %v, should give %d, but got %d", row, expect, res)
	}
}

func TestSample1(t *testing.T) {
	row := []int{0, 2, 1, 3}
	expect := 1
	runSample(t, row, expect)
}

func TestSample2(t *testing.T) {
	row := []int{3, 2, 0, 1}
	expect := 0
	runSample(t, row, expect)
}

func TestSample3(t *testing.T) {
	row := []int{5, 4, 2, 6, 3, 1, 0, 7}
	expect := 2
	runSample(t, row, expect)
}

func TestSample4(t *testing.T) {
	row := []int{6, 2, 1, 7, 4, 5, 3, 8, 0, 9}
	expect := 3
	runSample(t, row, expect)
}

func TestSample5(t *testing.T) {
	row := []int{5, 6, 4, 0, 2, 1, 9, 3, 8, 7, 11, 10}
	expect := 4
	runSample(t, row, expect)
}

func TestSample6(t *testing.T) {
	row := []int{5, 37, 38, 17, 28, 34, 42, 4, 22, 53, 57, 20, 23, 49, 3, 45, 24, 41, 10, 59, 48, 35, 30, 18, 15, 14, 2, 0, 33, 21, 36, 12, 31, 26, 39, 13, 9, 29, 27, 55, 1, 40, 19, 6, 11, 47, 7, 43, 44, 50, 16, 32, 52, 51, 56, 25, 54, 8, 46, 58}
	expect := 27
	runSample(t, row, expect)
}
