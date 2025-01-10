package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4 2
0 1 0 1
1 2
2 3
3 4
4 1
`
	expect := 50
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5 3
7 1 7 2
1 2
2 3
3 4
4 1
2 4
`
	expect := 96
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 30 10
372 372 298 372 372 657 298 298 657 298
1 3
1 6
1 7
1 8
1 9
1 10
2 3
2 6
2 7
2 10
3 4
3 5
3 6
3 9
4 6
4 7
4 8
4 9
4 10
5 6
5 7
5 8
5 9
5 10
6 7
6 8
6 10
7 9
8 9
9 10
`
	expect := 1045576
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 40 10
505 101 988 148 545 545 545 988 148 416
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
2 3
2 4
2 5
2 6
2 7
2 8
2 9
2 10
3 4
3 5
3 6
3 7
3 9
3 10
4 5
4 6
4 7
4 8
4 10
5 8
5 9
5 10
6 8
6 9
6 10
7 8
7 9
7 10
8 9
8 10
9 10
`
	expect := 1036928
	runSample(t, s, expect)
}
