package p2376

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := countSpecialNumbers(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 20
	expect := 19
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 135
	// 11, 22
	expect := 110
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 100
	// 11, 22, 33, 44, 55, 66, 77, 88, 99, 100
	expect := 90
	runSample(t, n, expect)
}
