package main

import "testing"

func runSample(t *testing.T, N int, C int, D int, E [][]int, S []int, expect int) {
	res, steps := solve(N, C, D, E, S)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var profit int

	chef := make([]int, N+1)

	chef[1] = C
	for _, step := range steps {
		if step == "nothing" {
			continue
		}
		if step[0] == 'b' {
			// build
			var u int
			readInt([]byte(step), 6, &u)
			if chef[u] == 0 {
				t.Fatalf("no chef at node %d", u)
			}
			profit += S[u-1]
		} else {
			// transfer a b c
			var u, v, c int
			pos := len("transfer") + 1
			pos = readInt([]byte(step), pos, &u)
			pos = readInt([]byte(step), pos+1, &v)
			readInt([]byte(step), pos+1, &c)
			if chef[u] < c {
				t.Fatalf("node %d has no sufficient chef to transfer to %d", u, v)
			}
			chef[v] += c
			chef[u] -= c
		}
	}

	if profit != expect {
		t.Fatalf("after all steps, no correct profit got, expect %d, but got %d", expect, profit)
	}
}

func TestSample1(t *testing.T) {
	N, C, D := 4, 2, 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	S := []int{-10, 5, 2, 6}
	expect := 11
	runSample(t, N, C, D, E, S, expect)
}
