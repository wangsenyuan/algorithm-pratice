package p1996

import "testing"

func runSample(t *testing.T, properties [][]int, expect int) {
	res := numberOfWeakCharacters(properties)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	props := [][]int{
		{5, 5}, {6, 3}, {3, 6},
	}
	expect := 0
	runSample(t, props, expect)
}

func TestSample2(t *testing.T) {
	props := [][]int{
		{2, 2}, {3, 3},
	}
	expect := 1
	runSample(t, props, expect)
}

func TestSample3(t *testing.T) {
	props := [][]int{
		{1, 5}, {10, 4}, {4, 3},
	}
	expect := 1
	runSample(t, props, expect)
}

func TestSample4(t *testing.T) {
	props := [][]int{
		{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10},
	}
	expect := 2
	runSample(t, props, expect)
}

func TestSample5(t *testing.T) {
	props := [][]int{
		{10, 1}, {5, 1}, {7, 10}, {4, 1}, {5, 9}, {6, 9}, {7, 2}, {1, 10},
	}
	expect := 4
	runSample(t, props, expect)
}
