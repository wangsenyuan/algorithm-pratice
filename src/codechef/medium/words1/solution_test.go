package main

import "testing"

func runSample(t *testing.T, n int, words []string, expect bool) {
	ws := make([][]byte, n)
	for i := 0; i < n; i++ {
		ws[i] = []byte(words[i])
	}

	res := solve(n, ws)

	if res != expect {
		t.Errorf("sample %d %v, expect %t, but got %t", n, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	words := []string{
		"skenzo",
		"logicboxes",
		"orderbox",
	}
	expect := true
	runSample(t, n, words, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	words := []string{
		"directi",
		"codechef",
	}
	expect := false
	runSample(t, n, words, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	words := []string{
		"ok",
		"ok",
	}
	expect := false
	runSample(t, n, words, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	words := []string{
		"ok",
		"ko",
	}
	expect := true
	runSample(t, n, words, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	words := []string{
		"bc", "de", "ab", "cd",
	}
	expect := true
	runSample(t, n, words, expect)
}

func TestSample6(t *testing.T) {
	n := 8
	words := []string{"ab", "ga", "bc", "cg", "fc", "cd", "ef", "de"}
	expect := true
	runSample(t, n, words, expect)
}
