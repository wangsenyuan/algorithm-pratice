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
	s := `8 5
0 1 2 9 3 2 7 5
2 2 1 9 4 1 5 8
2 6
1 7
2 4
7 8
5 8
`
	expect := []int{1, 3, 1, -1, -1}
	runSample(t, s, expect)
}
