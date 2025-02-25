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
	s := `5 3
2 1 4 3 5
1 2
3 5
1 4
`
	expect := []int{2, 0, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 10
2 1 5 3 4 6 9 8 7 10
3 9
2 5
4 8
5 6
3 8
2 10
7 8
6 7
8 10
4 10
`
	expect := []int{1, 3, 1, 2, 1, 0, 1, 1, 0, 0}
	runSample(t, s, expect)
}
