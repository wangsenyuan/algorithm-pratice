package p2594

import "testing"

func runSample(t *testing.T, ranks []int, cars int, expect int64) {
	res := repairCars(ranks, cars)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ranks := []int{4, 2, 3, 1}
	cars := 10
	var expect int64 = 16
	runSample(t, ranks, cars, expect)
}

func TestSample2(t *testing.T) {
	ranks := []int{5, 1, 8}
	cars := 6
	var expect int64 = 16
	runSample(t, ranks, cars, expect)
}

func TestSample3(t *testing.T) {
	ranks := []int{3}
	cars := 52
	var expect int64 = 8112
	runSample(t, ranks, cars, expect)
}
