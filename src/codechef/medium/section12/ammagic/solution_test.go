package main

import "testing"

func runSample(t *testing.T, tc, n, m int, B []string, expect int) {
	BB := make([][]byte, n)
	for i := 0; i < n; i++ {
		BB[i] = []byte(B[i])
	}
	res := solve(n, m, BB, tc)
	if res != expect {
		t.Errorf("Sample fail, expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	B := []string{
		"12", "34",
	}
	expect := 5
	runSample(t, 1, n, m, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 6
	B := []string{
		"020241",
		"103135",
		"861272",
		"529111",
	}
	expect := 22
	runSample(t, 2, n, m, B, expect)
}

func TestSample3(t *testing.T) {
	n, m := 10, 10
	B := []string{
		"9910778386",
		"2043114104",
		"3060911314",
		"7317085396",
		"1710280734",
		"5255219471",
		"9131541790",
		"4657660072",
		"6302126963",
		"1033147821",
	}
	expect := 129
	runSample(t, 3, n, m, B, expect)
}
