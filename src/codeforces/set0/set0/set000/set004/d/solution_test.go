package main

import "testing"

func runSample(t *testing.T, n, w, h int, envs [][]int, expect int) {
	res := solve(n, w, h, envs)

	if len(res) != expect {
		t.Errorf("Sample result %v, not as expect %d", res, expect)
		return
	}
	if expect == 0 {
		return
	}

	for i := 0; i < len(res); i++ {
		env := envs[res[i]-1]
		if env[0] <= w || env[1] <= h {
			t.Errorf("result %v, fail at check at step %d", res, i)
			return
		}
		w = env[0]
		h = env[1]
	}
}

func TestSample1(t *testing.T) {
	n, w, h := 2, 1, 1
	envs := [][]int{
		{2, 2},
		{2, 2},
	}
	expect := 1
	runSample(t, n, w, h, envs, expect)
}

func TestSample2(t *testing.T) {
	n, w, h := 3, 3, 3
	envs := [][]int{
		{5, 4},
		{12, 11},
		{9, 8},
	}
	expect := 3
	runSample(t, n, w, h, envs, expect)
}

func TestSample3(t *testing.T) {
	n, w, h := 4, 12, 140
	envs := [][]int{
		{172, 60},
		{71, 95},
		{125, 149},
		{53, 82},
	}
	expect := 1
	runSample(t, n, w, h, envs, expect)
}

func TestSample4(t *testing.T) {
	n, w, h := 15, 600, 875
	envs := [][]int{
		{1200, 451},
		{1664, 852},
		{1763, 1355},
		{1374, 1724},
		{1374, 1587},
		{1003, 1513},
		{1636, 1002},
		{431, 367},
		{1632, 690},
		{1257, 778},
		{410, 1632},
		{1045, 1279},
		{1762, 1763},
		{841, 576},
		{1165, 705},
	}
	expect := 3
	runSample(t, n, w, h, envs, expect)
}
