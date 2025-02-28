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
	s := `4 6 1
1 2 1
1 3 3
2 3 1
2 4 1
3 4 1
1 4 2
2
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 6 3
3 1 1
3 2 1
3 4 1
3 5 1
1 2 6
4 5 8
4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 38 8
2 1 10
3 2 782
4 2 678
5 3 967
6 3 291
7 1 64
8 2 995
9 3 876
10 1 579
6 10 658
6 2 931
3 10 81
8 3 245
2 9 527
3 4 739
6 1 909
5 2 951
5 6 132
6 9 359
7 4 928
7 6 656
9 4 786
8 10 155
1 9 201
9 10 35
6 8 884
2 7 648
4 6 709
7 9 780
8 4 855
3 1 663
9 8 984
9 5 774
1 5 18
2 10 152
8 7 588
8 5 360
7 3 584
893
`
	expect := 10
	runSample(t, s, expect)
}
