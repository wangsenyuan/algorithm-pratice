package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	a, res := process(reader)

	if len(res) > len(a) {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for _, op := range res {
		i, j := op[0]-1, op[1]-1
		a[i], a[j] = a[j], a[i]
		if abs(a[i]-a[j]) != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if !sort.IntsAreSorted(a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	s := `4
0 2 0 1`

	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 0`

	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6
0 1 1 2 2 2`

	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `4
2 2 1 0`

	// 2 1 2 0
	// 2 0 2 1
	// 1 0 2 2
	// 0 1 2 2

	runSample(t, s)
}
