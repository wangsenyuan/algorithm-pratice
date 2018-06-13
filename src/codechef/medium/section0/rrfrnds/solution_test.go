package main

import "testing"

func runSample(t *testing.T, n int, friends []string, expect int) {
	fs := make([][]byte, n)
	for i := 0; i < n; i++ {
		fs[i] = []byte(friends[i])
	}
	res := solve(n, fs)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, friends, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	fs := []string{
		"0111",
		"1000",
		"1000",
		"1000",
	}
	expect := 6
	runSample(t, n, fs, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	fs := []string{
		"0111",
		"1001",
		"1000",
		"1100",
	}
	expect := 4
	runSample(t, n, fs, expect)
}
