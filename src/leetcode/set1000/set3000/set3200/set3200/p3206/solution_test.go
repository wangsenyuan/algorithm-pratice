package p3206

import "testing"

func runSample(t *testing.T, a []int, current int, expect int) {
	res := maximumPoints(a, current)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 2}
	current := 2
	expect := 3
	runSample(t, a, current, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2}
	current := 10
	expect := 5
	runSample(t, a, current, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 2, 4, 1}
	current := 2
	expect := 13
	runSample(t, a, current, expect)
}
