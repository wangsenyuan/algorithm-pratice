package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	ok, res, mat := process(bufio.NewReader(strings.NewReader(s)))

	if expect < 0 {
		if ok {
			t.Errorf("Sample %s, expect %d, but got %t, %v", s, expect, ok, res)
		}
		return
	}
	if expect != len(res) || !ok {
		t.Errorf("Sample %s, expect %d, but got %t, %v", s, expect, ok, res)
	}

	for _, i := range res {
		i--
		mat[0][i], mat[1][i] = mat[1][i], mat[0][i]
	}

	check := func(arr []int) bool {
		sort.Ints(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] != i+1 {
				return false
			}
		}
		return true
	}

	if !check(mat[0]) || !check(mat[1]) {
		t.Fatalf("Sample result %v, not correct, leading to %v", res, mat)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 3 4
2 3 1 4`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
5 3 5 1 4
1 2 3 2 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2 1
3 3 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
1 2 2 1
3 4 3 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
4 3 1 4
3 2 2 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3
1 1 2
3 2 2`
	expect := -1
	runSample(t, s, expect)
}
