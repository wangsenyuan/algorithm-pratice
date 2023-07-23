package p2790

import "testing"

func runSample(t *testing.T, usageLimit []int, expect int) {
	res := maxIncreasingGroups(usageLimit)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	usageLimit := []int{1, 2, 5}
	expect := 3
	runSample(t, usageLimit, expect)
}

func TestSample2(t *testing.T) {
	usageLimit := []int{2, 1, 2}
	expect := 2
	runSample(t, usageLimit, expect)
}

func TestSample3(t *testing.T) {
	usageLimit := []int{1, 1}
	expect := 1
	runSample(t, usageLimit, expect)
}

func TestSample4(t *testing.T) {
	usageLimit := []int{7, 10, 4, 5, 7, 4, 5, 3, 5, 6}
	expect := 10
	runSample(t, usageLimit, expect)
}

func TestSample5(t *testing.T) {
	usageLimit := []int{2, 3}
	expect := 2
	runSample(t, usageLimit, expect)
}
