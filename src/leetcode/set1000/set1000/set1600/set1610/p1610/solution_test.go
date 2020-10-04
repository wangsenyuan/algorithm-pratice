package p1610

import "testing"

func runSample(t *testing.T, points [][]int, angle int, location []int, expect int) {
	res := visiblePoints(points, angle, location)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{2, 1}, {2, 2}, {3, 3}}
	angle := 90
	loc := []int{1, 1}
	expect := 3
	runSample(t, points, angle, loc, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}}
	angle := 90
	loc := []int{1, 1}
	expect := 4
	runSample(t, points, angle, loc, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{{0, 1}, {2, 1}}
	angle := 13
	loc := []int{1, 1}
	expect := 1
	runSample(t, points, angle, loc, expect)
}

func TestSample4(t *testing.T) {
	points := [][]int{{956, 232}, {438, 752}, {595, 297}, {508, 143}, {111, 594}, {645, 824}, {758, 434}, {447, 423}, {825, 356}, {807, 377}}
	angle := 38
	loc := []int{74, 581}
	expect := 7
	runSample(t, points, angle, loc, expect)
}
