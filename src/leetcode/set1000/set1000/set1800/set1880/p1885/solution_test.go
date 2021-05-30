package p1885

import "testing"

func runSample(t *testing.T, dist []int, speed int, hours int, expect int) {
	res := minSkips(dist, speed, hours)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dist := []int{1, 3, 2}
	speed := 4
	hours := 2
	expect := 1
	runSample(t, dist, speed, hours, expect)
}

func TestSample2(t *testing.T) {
	dist := []int{7, 3, 5, 5}
	speed := 2
	hours := 10
	expect := 2
	runSample(t, dist, speed, hours, expect)
}

func TestSample3(t *testing.T) {
	dist := []int{7, 3, 5, 5}
	speed := 1
	hours := 10
	expect := -1
	runSample(t, dist, speed, hours, expect)
}

func TestSample4(t *testing.T) {
	dist := []int{7, 6, 5, 3, 4, 6, 2, 2, 7, 2}
	speed := 5
	hours := 17
	expect := 0
	runSample(t, dist, speed, hours, expect)
}
