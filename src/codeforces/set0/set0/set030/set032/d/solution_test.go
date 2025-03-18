package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)
	if len(res) == 0 {
		if expect != "-1" {
			t.Fatalf("expect %s, but got %v", expect, res)
		}
		return
	}
	if expect == "-1" {
		t.Fatalf("expect -1, but got %v", res)
	}

	reader = bufio.NewReader(strings.NewReader(expect))
	for _, cur := range res {
		i, j := readTwoNums(reader)
		if i != cur[0] || j != cur[1] {
			t.Fatalf("expect %v, but got %v", cur, []int{i, j})
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 6 1
....*.
...***
....*.
..*...
.***..
`
	expect := `2 5
1 5
3 5
2 4
2 6
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 6 2
....*.
...***
....*.
..*...
.***..
`
	expect := `-1`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 7 2
...*...
.......
...*...
*.***.*
...*...
.......
...*...
`
	expect := `4 4
1 4
7 4
4 1
4 7
`
	runSample(t, s, expect)
}
