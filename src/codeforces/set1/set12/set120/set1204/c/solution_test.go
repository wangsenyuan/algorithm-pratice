package main

import "testing"

func runSample(t *testing.T, adj []string, p []int, expect []int) {
	q := make([]int, len(p))
	copy(q, p)

	res := solve(adj, q)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	var i, j int
	for i < len(res) {
		for j < len(p) && p[j] != res[i] {
			j++
		}
		if j == len(p) {
			t.Fatalf("Sample result %v, not a sub-sequence of %v", res, p)
		}
		j++
		i++
	}

	// res is a sub-sequence of p
	dist := getDist(adj)

	var sum int

	for i := 0; i+1 < len(p); i++ {
		u := p[i]
		v := p[i+1]
		if dist[u-1][v-1] < 0 {
			t.Fatalf("dist calculatiion for %d %d not correct", u, v)
		}
		sum += int(dist[u-1][v-1])
	}

	if sum+1 != len(p) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	adj := []string{
		"0110",
		"0010",
		"0001",
		"1000",
	}
	p := []int{1, 2, 3, 4}
	expect := []int{1, 2, 4}
	runSample(t, adj, p, expect)
}

func TestSample2(t *testing.T) {
	adj := []string{
		"0110",
		"0010",
		"1001",
		"1000",
	}
	p := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	expect := []int{1, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4}
	runSample(t, adj, p, expect)
}

func TestSample3(t *testing.T) {
	adj := []string{
		"011",
		"101",
		"110",
	}
	p := []int{1, 2, 3, 1, 3, 2, 1}
	expect := []int{1, 2, 3, 1, 3, 2, 1}
	runSample(t, adj, p, expect)
}

func TestSample4(t *testing.T) {
	adj := []string{
		"0110",
		"0001",
		"0001",
		"1000",
	}
	p := []int{1, 2, 4}
	expect := []int{1, 4}
	runSample(t, adj, p, expect)
}
