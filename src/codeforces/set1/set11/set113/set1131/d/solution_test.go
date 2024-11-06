package main

import "testing"

func runSample(t *testing.T, adj []string, expect int) {
	res, row, col := solve(adj)

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d %v %v", expect, res, row, col)
	}

	if expect < 0 {
		return
	}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(adj[i]); j++ {
			if adj[i][j] == '=' && row[i] != col[j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if adj[i][j] == '>' && row[i] <= col[j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if adj[i][j] == '<' && row[i] >= col[j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	adj := []string{
		">>>>",
		">>>>",
		">>>>",
	}
	expect := 2
	runSample(t, adj, expect)
}

func TestSample2(t *testing.T) {
	adj := []string{
		">>>",
		"<<<",
		">>>",
	}
	expect := 3
	runSample(t, adj, expect)
}

func TestSample3(t *testing.T) {
	adj := []string{
		"==",
		"=<",
		"==",
	}
	expect := -1
	runSample(t, adj, expect)
}
