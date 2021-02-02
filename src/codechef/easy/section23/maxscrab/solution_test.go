package main

import "testing"

func runSample(t *testing.T, S string, F []int, expect int64) {
	res := solve(len(S), []byte(S), F)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "..d.t.D..d"
	F := []int{10, 11, 12, 9, 8, 10, 11, 15}
	var expect int64 = 270
	runSample(t, S, F, expect)
}

func TestSample2(t *testing.T) {
	S := "dtDtTD..ddT.TtTdDT..TD"
	F := []int{12297, 5077, 28888, 17998, 12125, 27400, 31219, 21536}
	var expect int64 = 35629632
	runSample(t, S, F, expect)
}
