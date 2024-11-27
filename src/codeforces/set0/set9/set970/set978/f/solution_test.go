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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 2
10 4 10 15
1 2
4 3
`
	expect := []int{0, 0, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 4
5 4 1 5 4 3 7 1 2 5
4 6
2 1
10 8
3 5
`
	expect := []int{5, 4, 0, 5, 3, 3, 9, 0, 2, 5}
	runSample(t, s, expect)
}
