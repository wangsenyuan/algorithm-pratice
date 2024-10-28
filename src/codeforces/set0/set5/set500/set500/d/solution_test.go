package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, n int, roads [][]int, queries [][]int, expect []float64) {
	res := solve(n, roads, queries)

	check(t, expect, res)
}

func check(t *testing.T, expect []float64, res []float64) {

	for i, x := range expect {
		y := res[i]
		if math.Abs(x-y)/math.Max(1, x) >= 1e-7 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	data := `3
2 3 5
1 3 3
5
1 4
2 2
1 2
2 1
1 1
`
	reader := bufio.NewReader(strings.NewReader(data))
	res := process(reader)

	expect := []float64{
		14.0000000000, 12.0000000000, 8.0000000000, 6.0000000000, 4.0000000000,
	}

	check(t, expect, res)
}

func TestSample2(t *testing.T) {
	data := `6
1 5 3
5 3 2
6 1 7
1 4 4
5 2 3
5
1 2
2 1
3 5
4 1
5 2
`
	reader := bufio.NewReader(strings.NewReader(data))
	res := process(reader)

	expect := []float64{
		19.6000000000, 18.6000000000, 16.6000000000, 13.6000000000, 12.6000000000,
	}

	check(t, expect, res)
}
