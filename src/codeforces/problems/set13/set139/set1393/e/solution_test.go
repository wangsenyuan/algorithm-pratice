package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []string{
		"abcd",
		"zaza",
		"ataka",
	}
	expect := 4
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := []string{
		"dfs",
		"bfs",
		"sms",
		"mms",
	}
	expect := 8
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	S := []string{
		"abc",
		"bcd",
		"a",
	}
	expect := 0
	runSample(t, n, S, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	S := []string{
		"lapochka",
		"kartyshka",
		"bigbabytape",
		"morgenshtern",
		"ssshhhiiittt",
		"queen",
	}
	expect := 2028
	runSample(t, n, S, expect)
}
