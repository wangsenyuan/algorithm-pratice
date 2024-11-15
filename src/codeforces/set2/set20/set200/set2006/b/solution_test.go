package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1000000000000
1
2 1000000000000`
	expect := []int{2000000000000}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 9
1 1 1
2 2
4 4
3 3`
	expect := []int{25, 18, 18}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 100
1 2 3 2 1
6 17
3 32
2 4
4 26
5 21`
	expect := []int{449, 302, 247, 200, 200}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 511
1 2 2 4 2 1 1 8 8
3 2
6 16
10 256
9 128
2 1
5 8
8 64
4 4
7 32`
	expect := []int{4585, 4473, 2681, 1567, 1454, 1322, 1094, 1022, 1022}
	runSample(t, s, expect)
}
