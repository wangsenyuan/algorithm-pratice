package p2332

import "testing"

func runSample(t *testing.T, buses []int, passengers []int, capacity int, expect int) {
	res := latestTimeCatchTheBus(buses, passengers, capacity)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	buses := []int{3}
	passengers := []int{2, 3}
	capacity := 2
	expect := 1
	runSample(t, buses, passengers, capacity, expect)
}

func TestSample2(t *testing.T) {
	buses := []int{2}
	passengers := []int{2}
	capacity := 2
	expect := 1
	runSample(t, buses, passengers, capacity, expect)
}

func TestSample3(t *testing.T) {
	buses := []int{5, 15}
	passengers := []int{11, 12, 13, 14, 15}
	capacity := 2
	expect := 10
	runSample(t, buses, passengers, capacity, expect)
}

func TestSample4(t *testing.T) {
	buses := []int{3}
	passengers := []int{4}
	capacity := 1
	expect := 3
	runSample(t, buses, passengers, capacity, expect)
}

func TestSample5(t *testing.T) {
	buses := []int{15, 16, 17, 7, 10, 20, 13, 12}
	passengers := []int{18, 15, 11, 17, 12, 13, 14, 10, 19, 16}
	capacity := 2
	expect := 9
	runSample(t, buses, passengers, capacity, expect)
}

func TestSample6(t *testing.T) {
	buses := []int{6, 8, 18, 17}
	passengers := []int{6, 8, 17}
	capacity := 1
	expect := 18
	runSample(t, buses, passengers, capacity, expect)
}
