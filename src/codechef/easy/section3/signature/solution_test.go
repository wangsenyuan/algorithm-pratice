package main

import "testing"

func runSample(t *testing.T, n, m int, A, B []string, expect int) {
	res := solve(n, m, A, B)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	A := []string{
		"100",
		"010",
		"000",
	}
	B := []string{
		"000",
		"010",
		"001",
	}
	runSample(t, n, m, A, B, 0)
}

func TestSample2(t *testing.T) {
	n, m := 4, 4
	A := []string{
		"0000",
		"0110",
		"0000",
		"0011",
	}
	B := []string{
		"1100",
		"0000",
		"1100",
		"0000",
	}
	runSample(t, n, m, A, B, 2)
}

func TestSample4(t *testing.T) {
	n, m := 3, 3
	A := []string{
		"111",
		"000",
		"000",
	}
	B := []string{
		"001",
		"001",
		"001",
	}
	runSample(t, n, m, A, B, 2)
}
