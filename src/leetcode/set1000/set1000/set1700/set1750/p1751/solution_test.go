package p1751

import "testing"

func runSample(t *testing.T, events [][]int, k int, expect int) {
	res := maxValue(events, k)
	if res != expect {
		t.Errorf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	events := [][]int{
		{1, 2, 4}, {3, 4, 3}, {2, 3, 1},
	}
	k := 2
	expect := 7
	runSample(t, events, k, expect)
}

func TestSample2(t *testing.T) {
	events := [][]int{
		{1, 2, 4}, {3, 4, 3}, {2, 3, 10},
	}
	k := 2
	expect := 10
	runSample(t, events, k, expect)
}

func TestSample3(t *testing.T) {
	events := [][]int{
		{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4},
	}
	k := 3
	expect := 9
	runSample(t, events, k, expect)
}

func TestSample4(t *testing.T) {
	events := [][]int{
		{19, 42, 7}, {41, 73, 15}, {52, 73, 84}, {84, 92, 96}, {6, 64, 50}, {12, 56, 27}, {22, 74, 44}, {38, 85, 61},
	}
	k := 5
	expect := 187
	runSample(t, events, k, expect)
}
