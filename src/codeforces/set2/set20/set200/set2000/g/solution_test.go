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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
100 20 80
1 5 30 100
1 2 20 50
2 3 20 50
3 4 20 50
4 5 20 50`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
100 50 60
1 2 55 110`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
100 40 60
1 2 30 100
2 4 30 100
1 3 20 50
3 4 20 50`
	expect := 60
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3
100 80 90
1 2 1 10
2 3 10 50
1 3 20 21`
	expect := 80
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 2
58 55 57
2 1 1 3
2 3 3 4`
	expect := 53
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 1
12 9 10
2 1 6 10`
	expect := 3
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `5 5
8 5 6
2 1 1 8
2 3 4 8
4 2 2 4
5 3 3 4
4 5 2 6`
	expect := 2
	runSample(t, s, expect)
}
