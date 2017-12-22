package main

import "testing"

func runSample(id int, t *testing.T, n int, flights [][]int, from int, begin int, end int, birthday int, can bool, cnt int) {
	rcan, rcnt := solve(n, flights, from, begin, end, birthday)
	if can && !rcan {
		t.Errorf("the sample %d should be able to reach the poly-poly, but got not", id)
	} else if !can && rcan {
		t.Errorf("the sample %d should not be able to reach the party, but got yes", id)
	} else if can && cnt != rcnt {
		t.Errorf("the sample %d should be able to reach the party after %d flights, but got %d", id, cnt, rcnt)
	}
}

func TestSample1(t *testing.T) {
	flights := [][]int{
		{1, 3, 4, 3},
		{4, 4, 3, 5},
		{3, 10, 2, -1},
	}
	runSample(1, t, 3, flights, 1, 2, 2, 0, true, 3)
}

func TestSample2(t *testing.T) {
	flights := [][]int{
		{1, 2, 3, 0},
	}
	runSample(2, t, 1, flights, 1, 2, 2, 2, false, -1)
}

func TestSample3(t *testing.T) {
	flights := [][]int{
		{2, 0, 2, 0},
		{2, 1, 4, 0},
	}
	runSample(3, t, 2, flights, 2, 0, 4, 0, false, -1)
}
