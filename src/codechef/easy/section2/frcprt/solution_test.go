package main

import "testing"

func runSample(t *testing.T, N, M int, A []string, S string, expect []string) {
	B := make([][]byte, N)
	for i := 0; i < N; i++ {
		B[i] = []byte(A[i])
	}
	res := solve(N, M, B, []byte(S))
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if expect[i][j] != res[i][j] {
				t.Errorf("Sample %d %d %v %s, expect %v, but got wrong answer at %d %d", N, M, A, S, expect, i, j)
				return
			}
		}
	}
}

func TestSample1(t *testing.T) {
	N, M := 4, 4
	A := []string{
		"1010",
		"0010",
		"1001",
		"0100",
	}
	S := "LRDU"
	expect := []string{
		"0011",
		"0011",
		"0001",
		"0001",
	}
	runSample(t, N, M, A, S, expect)
}

func TestSample2(t *testing.T) {
	N, M := 4, 3
	A := []string{
		"000",
		"010",
		"001",
		"101",
	}
	S := "LRL"
	expect := []string{
		"000",
		"100",
		"100",
		"110",
	}
	runSample(t, N, M, A, S, expect)
}

func TestSample3(t *testing.T) {
	N, M := 3, 2
	A := []string{
		"01",
		"10",
		"00",
	}
	S := "D"
	expect := []string{
		"00",
		"00",
		"11",
	}
	runSample(t, N, M, A, S, expect)
}
