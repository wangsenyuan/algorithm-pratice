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
	s := `4
1 2
1 3
4 1
0101`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2
3 2
2 4
???0`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2
1 3
2 4
2 5
?1?01`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 2
2 3
3 4
5 3
3 6
?0????`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
1 2
1 3
1 4
1 5
11?1?`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2
2 1
??`
	expect := 0
	runSample(t, s, expect)
}
