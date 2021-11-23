package p2079

import "testing"

func runSample(t *testing.T, plants []int, capacity int, expect int) {
	res := wateringPlants(plants, capacity)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	plants := []int{2, 2, 3, 3}
	capacity := 5
	expect := 14
	runSample(t, plants, capacity, expect)
}

func TestSample2(t *testing.T) {
	plants := []int{1, 1, 1, 4, 2, 3}
	capacity := 4
	expect := 30
	runSample(t, plants, capacity, expect)
}

func TestSample3(t *testing.T) {
	plants := []int{7, 7, 7, 7, 7, 7, 7}
	capacity := 8
	expect := 49
	runSample(t, plants, capacity, expect)
}
