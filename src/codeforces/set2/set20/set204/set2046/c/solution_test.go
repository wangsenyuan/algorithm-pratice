package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	cnt, res, points := process(bufio.NewReader(strings.NewReader(s)))

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d, %v %v", expect, cnt, res, points)
	}

	sum := make([]int, 4)
	x0, y0 := res[0], res[1]
	for _, cur := range points {
		x, y := cur[0], cur[1]
		if x0 <= x && y0 <= y {
			sum[0]++
		} else if x0 > x && y0 <= y {
			sum[1]++
		} else if x0 <= x && y0 > y {
			sum[2]++
		} else {
			sum[3]++
		}
	}

	w := slices.Min(sum)

	if w != cnt {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 1
1 2
2 1
2 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
0 0
0 0
0 0
0 0`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
1 2
2 1
2 -1
1 -2
-1 -2
-2 -1
-2 1
-1 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
1 1
1 2
1 3
1 4
2 1
3 1
4 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `15
1 -3
9 -2
-4 -8
-9 5
2 6
-10 -5
9 -2
6 0
4 -7
-2 -5
7 7
7 7
-2 -9
9 7
2 10`
	expect := 3
	runSample(t, s, expect)
}
