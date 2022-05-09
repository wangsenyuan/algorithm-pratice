package main

import "testing"

func runSample(t *testing.T, n int, m int, A []string, S string, expect bool) {
	res := solve(n, m, A, S)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var pos int

	for _, cur := range res {
		l, r, i := cur[0], cur[1], cur[2]
		l--
		r--
		i--
		next := pos + (r - l + 1)
		if A[i][l:r+1] != S[pos:next] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		pos = next
	}

	if pos != m {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 8
	A := []string{
		"12340219",
		"20215601",
		"56782022",
		"12300678",
	}
	S := "12345678"
	expect := true
	runSample(t, n, m, A, S, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 3
	A := []string{
		"134",
		"126",
	}
	S := "123"
	expect := false
	runSample(t, n, m, A, S, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 4
	A := []string{
		"1210",
	}
	S := "1221"
	expect := true
	runSample(t, n, m, A, S, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 3
	A := []string{
		"251",
		"064",
		"859",
		"957",
	}
	S := "054"
	expect := false
	runSample(t, n, m, A, S, expect)
}

func TestSample5(t *testing.T) {
	n, m := 4, 7
	A := []string{
		"7968636",
		"9486033",
		"4614224",
		"5454197",
	}
	S := "9482268"
	expect := true
	runSample(t, n, m, A, S, expect)
}
