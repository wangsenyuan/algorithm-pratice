package p1893

import "testing"

func runSample(t *testing.T, ranges [][]int, left, right int, expect bool) {
	res := isCovered(ranges, left, right)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ranges := [][]int{{1, 2}, {3, 4}, {5, 6}}
	left := 2
	right := 5
	expect := true
	runSample(t, ranges, left, right, expect)
}

func TestSample2(t *testing.T) {
	ranges := [][]int{{1, 10}, {10, 20}}
	left := 21
	right := 21
	expect := false
	runSample(t, ranges, left, right, expect)
}

func TestSample3(t *testing.T) {
	ranges := [][]int{{1, 1}, {3, 3}}
	left := 3
	right := 3
	expect := true
	runSample(t, ranges, left, right, expect)
}

func TestSample4(t *testing.T) {
	ranges := [][]int{{1, 1}, {20, 20}}
	left := 21
	right := 21
	expect := false
	runSample(t, ranges, left, right, expect)
}
