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
	s := `1 4
6 3`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 5
4 4`
	//这个例子很难理解啊
	//
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 10
1 1
1 1
1 1
1 1
1 1`

	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 100
1 2
5 6`

	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 11
2 2
3 3
4 4`

	expect := 17
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 25
9 2
4 3
8 10`

	expect := 80
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4 18
5 4
8 5
8 3
6 2`

	expect := 35
	runSample(t, s, expect)
}
