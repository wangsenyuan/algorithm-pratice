package b

import "testing"

func runSample(t *testing.T, target []int, expect bool) {
	res := isMagic(target)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := []int{2, 4, 3, 1, 5}
	expect := true
	runSample(t, target, expect)
}

func TestSample2(t *testing.T) {
	target := []int{5, 4, 3, 2, 1}
	expect := false
	runSample(t, target, expect)
}

func TestSample3(t *testing.T) {
	target := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 55, 59, 63, 67, 53, 57, 61, 65, 69}
	expect := true
	runSample(t, target, expect)
}

func TestSample4(t *testing.T) {
	target := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 35, 41, 33, 39, 37, 47, 43, 45}
	expect := false
	runSample(t, target, expect)
}
