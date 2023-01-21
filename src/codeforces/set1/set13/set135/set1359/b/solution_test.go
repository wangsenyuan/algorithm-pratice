package main

import "testing"

func runSample(t *testing.T, n, m int, M []string, x, y int, expect int64) {
	MM := make([][]byte, n)
	for i := 0; i < n; i++ {
		MM[i] = []byte(M[i])
	}
	res := solve(n, m, MM, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, x, y := 1, 1, 10, 1
	M := []string{"."}
	var expect int64 = 10
	runSample(t, n, m, M, x, y, expect)
}

func TestSample2(t *testing.T) {
	n, m, x, y := 1, 2, 10, 1
	M := []string{".."}
	var expect int64 = 1
	runSample(t, n, m, M, x, y, expect)
}

func TestSample3(t *testing.T) {
	n, m, x, y := 2, 1, 10, 1
	M := []string{".", "."}
	var expect int64 = 20
	runSample(t, n, m, M, x, y, expect)
}

func TestSample4(t *testing.T) {
	n, m, x, y := 3, 3, 3, 7
	M := []string{
		"..*",
		"*..",
		".*.",
	}
	var expect int64 = 18
	runSample(t, n, m, M, x, y, expect)
}
