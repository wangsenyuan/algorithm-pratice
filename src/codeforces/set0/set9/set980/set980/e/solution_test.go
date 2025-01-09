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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
2 1
2 6
4 2
5 6
2 3
`
	expect := []int{1, 3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 4
2 6
2 7
7 8
1 2
3 1
2 4
7 5
`
	expect := []int{1, 3, 4, 5}
	runSample(t, s, expect)
}
