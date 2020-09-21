package p1596

import "testing"

func runSample(t *testing.T, cost [][]int, expect int) {
	res := connectTwoGroups(cost)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", cost, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cost := [][]int{
		{15, 96}, {36, 2},
	}
	expect := 17
	runSample(t, cost, expect)
}

func TestSample2(t *testing.T) {
	cost := [][]int{
		{1, 3, 5}, {4, 1, 1}, {1, 5, 3},
	}
	expect := 4
	runSample(t, cost, expect)
}

func TestSample3(t *testing.T) {
	cost := [][]int{
		{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8},
	}
	expect := 10
	runSample(t, cost, expect)
}

func TestSample4(t *testing.T) {
	cost := [][]int{
		{93, 56, 92}, {53, 44, 18}, {86, 44, 69}, {54, 60, 30},
	}
	expect := 172
	runSample(t, cost, expect)
}
