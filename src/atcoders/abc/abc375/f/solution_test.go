package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 5
1 2 5
1 3 10
2 3 6
2 1 3
1 2
2 1 3
1 1
2 1 3
`
	expect := []int64{10, 11, -1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6 6
2 3 1
2 4 1
3 4 1
1 2 1
1 3 1
1 4 1
1 4
1 5
1 6
2 1 2
2 1 3
2 1 4
`
	expect := []int64{-1, -1, -1}
	runSample(t, s, expect)
}
