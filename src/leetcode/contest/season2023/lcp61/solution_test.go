package lcp61

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := temperatureTrend(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, -15, 3, 14, -1, 4, 35, 36}
	B := []int{-15, 32, 20, 9, 33, 4, -1, -5}
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-14, 7, -19, 9, 13, 40, 19, 15, -18}
	B := []int{3, 16, 28, 32, 25, 12, 13, -6, 4}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{21, 18, 18, 18, 31}
	B := []int{34, 32, 16, 16, 17}
	expect := 2
	runSample(t, A, B, expect)
}
