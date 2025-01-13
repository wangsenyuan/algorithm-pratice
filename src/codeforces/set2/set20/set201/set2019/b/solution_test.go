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
	n := len(res)

	arr := readNNums(bufio.NewReader(strings.NewReader(expect)), n)

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
101 200
2 1`
	expect := "0 100"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 15
1 2 3 5 6 7
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15`
	expect := "0 0 0 0 2 0 0 0 3 0 2 0 0 0 0"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 8
254618033 265675151 461318786 557391198 848083778
6 9 15 10 6 9 4 4294967300`
	expect := "291716045 0 0 0 291716045 0 301749698 0"
	runSample(t, s, expect)
}
