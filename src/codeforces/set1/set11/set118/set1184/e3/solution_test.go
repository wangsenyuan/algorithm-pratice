package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 8
2 3 3
3 1 4
`
	expect := []int{4, 8, 8}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 9
5 10 606854707
3 8 737506631
2 4 429066157
8 9 947792932
6 4 56831480
2 5 541638168
10 7 20498997
7 9 250445792
6 1 9522145
`
	expect := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 6
1 2 1
2 3 1
1 4 1
4 5 1
3 5 3
2 4 4
`
	expect := []int{3, 3, 3, 3, 1, 1}
	runSample(t, s, expect)
}
