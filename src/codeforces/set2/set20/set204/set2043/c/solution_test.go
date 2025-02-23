package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	reader = bufio.NewReader(strings.NewReader(expect))
	n := readNum(reader)
	ans := readNNums(reader, n)
	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 -1 10 1 1`
	expect := `8
-1 0 1 2 9 10 11 12 `

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
-1 -1 -1 -1 -1`
	expect := `6
-5 -4 -3 -2 -1 0 `

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
-1 2`
	expect := `4
-1 0 1 2`

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
7 1`
	expect := `4
0 1 7 8 `

	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
1 4 -1`
	expect := `6
-1 0 1 3 4 5`

	runSample(t, s, expect)
}
