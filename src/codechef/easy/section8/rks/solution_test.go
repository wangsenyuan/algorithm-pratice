package main

import "testing"

func runSample(t *testing.T, N, K int, rks [][]int, p int) {
	res := solve(N, K, rks)
	if len(res) != p {
		t.Fatalf("sample %d %d %v, expect %d, but got %d", N, K, rks, p, len(res))
	}

	row := make([]int, N)
	col := make([]int, N)
	for _, rk := range rks {
		x, y := rk[0], rk[1]
		row[x-1]++
		col[y-1]++
	}

	for _, rk := range res {
		x, y := rk[0], rk[1]
		if row[x-1] > 0 {
			t.Errorf("more than one rooks in row %d", x)
		}
		row[x-1]++
		if col[y-1] > 0 {
			t.Errorf("more than one rooks at col %d", y)
		}
		col[y-1]++
	}
}

func TestSample1(t *testing.T) {
	N, K := 4, 2
	rks := [][]int{{1, 4}, {2, 2}}
	runSample(t, N, K, rks, 2)
}

func TestSample2(t *testing.T) {
	N, K := 4, 0
	runSample(t, N, K, nil, 4)
}
