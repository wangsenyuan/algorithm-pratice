package main

import "testing"

func runSample(t *testing.T, s string, expect_cost int, expect_length int) {
	cost, res := solve(s)

	if cost != expect_cost || len(res) != expect_length {
		t.Fatalf("Sample expect %d, %d, but got %d, %v", expect_cost, expect_length, cost, res)
	}

	var tmp int
	for i := 1; i < len(res); i++ {
		j := res[i] - 1
		k := res[i-1] - 1
		tmp += abs(index(s[j]) - index(s[k]))
	}

	if tmp != expect_cost {
		t.Fatalf("Sample result %v, get wrong cost %d", res, tmp)
	}
}

func TestSample1(t *testing.T) {
	s := "logic"
	cost := 9
	length := 4
	runSample(t, s, cost, length)
}
