package main

import "testing"

func runSample(t *testing.T, L, R, K int64, N int, S []string, expect string) {
	res := solve(L, R, K, N, S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var L, R, K int64 = 1, 1000000000, 4
	S := []string{
		"62", "63",
	}
	expect := "163"
	runSample(t, L, R, K, len(S), S, expect)
}

func TestSample2(t *testing.T) {
	var L, R, K int64 = 1, 1, 1
	S := []string{
		"2",
	}
	expect := "no such number"
	runSample(t, L, R, K, len(S), S, expect)
}

func TestSample3(t *testing.T) {
	var L, R, K int64 = 1, 1000, 15
	S := []string{
		"6", "22",
	}
	expect := "67"
	runSample(t, L, R, K, len(S), S, expect)
}
