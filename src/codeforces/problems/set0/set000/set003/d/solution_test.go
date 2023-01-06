package main

import "testing"

func runSample(t *testing.T, s string, cost [][]int, expect int64) {
	res, ss := solve(s, cost)

	if res != expect {
		t.Fatalf("Sample %s %v, expect %d, but got %d", s, cost, expect, res)
	}
	var ans int64
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] == '?' {
			if ss[i] == '(' {
				ans += int64(cost[j][0])
			} else {
				ans += int64(cost[j][1])
			}
			j++
		}
	}
	if ans != expect {
		t.Fatalf("Sample result %s, not correct", ss)
	}
	var open int
	for i := 0; i < len(ss); i++ {
		if ss[i] == '(' {
			open++
		} else if ss[i] == ')' {
			open--
		}
		if open < 0 {
			t.Fatalf("Sample result %s, not correct", ss)
		}
	}
	if open != 0 {
		t.Fatalf("Sample result %s, not correct", ss)
	}
}

func TestSample1(t *testing.T) {
	s := "(??)"
	cost := [][]int{
		{1, 2}, {2, 8},
	}
	var expect int64 = 4
	runSample(t, s, cost, expect)
}
