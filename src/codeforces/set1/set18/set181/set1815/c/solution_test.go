package main

import "testing"

func runSample(t *testing.T, n int, requirements [][]int, expect int) {
	res := solve(n, requirements)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if expect == 0 {
		return
	}
	var cnt int
	for _, x := range res {
		if x <= 1 || x > n {
			cnt++
		}
	}
	if cnt != 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for _, cur := range requirements {
		a, b := cur[0], cur[1]
		if a == 1 {
			continue
		}
		// 所有的a中间都必须要有b
		var buf []int
		for _, x := range res {
			if x == a || x == b {
				buf = append(buf, x)
			}
		}
		if len(buf)%2 != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		// a,b,a,b,a
		for i := 0; i < len(buf); i++ {
			if i%2 == 0 && buf[i] != a {
				t.Fatalf("Sample result %v, not correct", res)
			} else if i%2 == 1 && buf[i] != b {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	req := [][]int{
		{3, 1},
		{2, 1},
	}
	expect := 5
	runSample(t, n, req, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	expect := 1
	runSample(t, n, nil, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	expect := 0
	runSample(t, n, nil, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	reqs := [][]int{
		{1, 2},
		{2, 1},
	}
	expect := 3
	runSample(t, n, reqs, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	reqs := [][]int{
		{2, 1},
		{3, 1},
		{4, 2},
		{4, 5},
		{5, 1},
	}
	expect := 10
	runSample(t, n, reqs, expect)
}
