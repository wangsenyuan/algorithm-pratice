package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 10, 7)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 5, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, 9, 11, 11)
}

func TestSampleData(t *testing.T) {
	in := `55
1 1
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
2 2
2 3
2 4
2 5
2 6
2 7
2 8
2 9
2 10
3 3
3 4
3 5
3 6
3 7
3 8
3 9
3 10
4 4
4 5
4 6
4 7
4 8
4 9
4 10
5 5
5 6
5 7
5 8
5 9
5 10
6 6
6 7
6 8
6 9
6 10
7 7
7 8
7 9
7 10
8 8
8 9
8 10
9 9
9 10
10 10
	`
	reader := strings.NewReader(in)

	expect := `1
1
3
3
3
3
7
7
7
7
2
3
3
3
3
7
7
7
7
3
3
3
3
7
7
7
7
4
5
5
7
7
7
7
5
5
7
7
7
7
6
7
7
7
7
7
7
7
7
8
9
9
9
9
10
`

	res := process(bufio.NewReader(reader))

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}
