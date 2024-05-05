package main

import "testing"

func runSample(t *testing.T, s []int, expect bool) {
	res := solve(s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample %v, expect %t, but got %v", s, expect, res)
	}
	if !expect {
		return
	}
	rec := make(map[int]int)
	for i := 0; i < len(res); i++ {
		var g int
		for j := i; j < len(res); j++ {
			g = gcd(g, res[j])
			rec[g]++
		}
	}

	if len(rec) != len(s) {
		t.Fatalf("Sample result %v, not correct", res)
	}
	for _, x := range s {
		if rec[x] == 0 {
			t.Fatalf("Sample result %v, which was not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := []int{2, 4, 6, 12}
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []int{2, 3}
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []int{1, 2, 3, 4, 6, 12}
	expect := true
	runSample(t, s, expect)
}
