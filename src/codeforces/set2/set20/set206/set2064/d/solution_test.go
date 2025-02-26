package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	ans := readNNums(bufio.NewReader(strings.NewReader(expect)), len(res))

	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
5
6`
	expect := "1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
1 5 4 11
8
13
16
15`
	expect := "0 2 4 2"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 9
10 4 3 9 7 4 6 1 9 4
2
6
5
6
9
8
6
2
7`
	expect := "0 1 1 1 3 3 1 0 1"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 4
6 17 54 40 45 18 14 7
13
61
17
3`
	expect := "1 3 3 0"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 5
48 11 60 18 60 12
22
22
27
56
35`
	expect := "1 1 1 1 1 "
	runSample(t, s, expect)
}
