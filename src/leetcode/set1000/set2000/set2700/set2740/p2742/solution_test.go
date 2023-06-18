package p2742

import "testing"

func runSample(t *testing.T, cost []int, time []int, expect int) {
	res := paintWalls(cost, time)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cost := []int{1, 2, 3, 2}
	time := []int{1, 2, 3, 2}
	expect := 3
	runSample(t, cost, time, expect)
}

func TestSample2(t *testing.T) {
	cost := []int{2, 3, 4, 2}
	time := []int{1, 1, 1, 1}
	expect := 4
	runSample(t, cost, time, expect)
}

func TestSample3(t *testing.T) {
	cost := []int{26, 53, 10, 24, 25, 20, 63, 51}
	time := []int{1, 1, 1, 1, 2, 2, 2, 1}
	expect := 55
	runSample(t, cost, time, expect)
}
