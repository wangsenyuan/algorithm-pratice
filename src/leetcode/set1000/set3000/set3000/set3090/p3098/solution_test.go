package p3098

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := minimumDistance(points)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}
	expect := 12
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expect := 0
	runSample(t, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{{3, 2}, {3, 9}, {7, 10}, {4, 4}, {8, 10}, {2, 7}}
	expect := 10
	runSample(t, points, expect)
}
