package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for i, x := range expect {
		y := res[i]
		if math.Abs(x-y) > 1e-9 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7
-1 8
1 4
1 12
2 2
2 6
3 10
3 14
1
1
`
	expect := []float64{8.0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
-1 5
1 3
1 7
6
1
2
4
6
8
9
`
	expect := []float64{7.0, 7, 7, 3, 3, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 5
-1 10
1 3
1 7
2 15
6
1
4
6
8
11
20
`
	expect := []float64{11.0000000000, 11.0000000000, 9.0000000000, 9.0000000000, 7.0000000000, 7.0000000000}
	runSample(t, s, expect)
}
