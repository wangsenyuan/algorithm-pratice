package p871

import "testing"

func runSampe(t *testing.T, target int, startFuel int, stations [][]int, expect int) {
	res := minRefuelStops(target, startFuel, stations)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", target, startFuel, stations, expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := 1
	startFuel := 1
	stations := [][]int{}
	expect := 0
	runSampe(t, target, startFuel, stations, expect)
}

func TestSample2(t *testing.T) {
	target := 100
	startFuel := 1
	stations := [][]int{{10, 100}}
	expect := -1
	runSampe(t, target, startFuel, stations, expect)
}

func TestSample3(t *testing.T) {
	target := 100
	startFuel := 10
	stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	expect := 2
	runSampe(t, target, startFuel, stations, expect)
}
