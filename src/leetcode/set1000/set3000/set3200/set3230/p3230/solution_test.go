package p3230

import "testing"

func runSample(t *testing.T, x int, y int, circles [][]int, expect bool) {
	res := canReachCorner(x, y, circles)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 3, 4
	circles := [][]int{
		{2, 1, 1},
	}
	expect := true
	runSample(t, x, y, circles, expect)
}

func TestSample2(t *testing.T) {
	x, y := 3, 3
	circles := [][]int{
		{1, 1, 2},
	}
	expect := false
	runSample(t, x, y, circles, expect)
}

func TestSample3(t *testing.T) {
	x, y := 3, 3
	circles := [][]int{
		{2, 1, 1}, {1, 2, 1},
	}
	expect := false
	runSample(t, x, y, circles, expect)
}
