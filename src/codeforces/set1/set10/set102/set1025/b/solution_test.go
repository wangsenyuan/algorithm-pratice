package main

import "testing"

func runSample(t *testing.T, pairs [][]int, expect int) {
	res := solve(pairs)

	if (expect > 0) != (res > 0) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	if expect < 0 {
		return
	}

	for _, cur := range pairs {
		if cur[0]%res != 0 && cur[1]%res != 0 {
			t.Fatalf("Sample result %d, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{17, 18},
		{15, 24},
		{12, 15},
	}
	expect := 6
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{10, 16},
		{7, 17},
	}
	expect := -1
	runSample(t, pairs, expect)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{
		{893949877, 894069401},
		{893949877, 894069401},
	}
	expect := 893949877
	runSample(t, pairs, expect)
}

func TestSample4(t *testing.T) {
	pairs := [][]int{
		{1435781229, 1041192184},
		{253359306, 340480943},
		{1872995232, 28582653},
		{1337658848, 1307583218},
		{1637848533, 535648393},
		{486937247, 1782881418},
		{871715693, 364823034},
		{1270203507, 689669064},
		{49503294, 1049495360},
		{1239477847, 87721143},
		{1053001443, 747987815},
		{1918849059, 1955601974},
		{160018454, 936007735},
		{1559803677, 246706454},
		{1955942252, 1606295975},
		{1072004390, 429250409},
		{1620918269, 894807105},
		{682295613, 325414994},
		{1881594207, 284980342},
		{1568966745, 351921835},
		{184615154, 1786427452},
	}
	expect := 893949877
	runSample(t, pairs, expect)
}
