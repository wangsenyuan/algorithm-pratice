package main

import "testing"

func runSample(t *testing.T, k int, tasks [][]int, expect int) {
	profit, L, R, set := solve(k, tasks)

	if profit != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, profit)
	}
	if expect == 0 {
		return
	}

	res := -(R - L + 1) * k

	for _, i := range set {
		l, r, p := tasks[i-1][0], tasks[i-1][1], tasks[i-1][2]
		if L <= l && r <= R {
			res += p
		} else {
			t.Fatalf("Sample result %v, not valid", set)
		}
	}

	if res != expect {
		t.Errorf("Sample result %v, gettting profit %d, but not expected %d", set, res, expect)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	tasks := [][]int{
		{1, 1, 3},
		{3, 3, 11},
		{5, 5, 17},
		{7, 7, 4},
	}
	expect := 13
	runSample(t, k, tasks, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	tasks := [][]int{
		{1, 2, 5},
	}
	expect := 0
	runSample(t, k, tasks, expect)
}

func TestSample3(t *testing.T) {
	k := 8
	tasks := [][]int{
		{1, 5, 16},
		{2, 4, 9},
		{3, 3, 24},
		{1, 5, 13},
	}
	expect := 22
	runSample(t, k, tasks, expect)
}

func TestSample4(t *testing.T) {
	k := 666618512
	tasks := [][]int{
		{1, 1, 751101774540 + 666618512},
	}
	for i := 2; i <= 10000; i++ {
		tasks = append(tasks, []int{i, i, 666618512})
	}
	expect := 751101774540
	runSample(t, k, tasks, expect)
}
