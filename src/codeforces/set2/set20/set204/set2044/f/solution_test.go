package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 6
-2 3 -3
-2 2 -1
-1
1
-2
2
-3
3
`
	expect := []bool{false, true, false, false, true, false}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5 6
1 -2 3 0 0
0 -2 5 0 -3
4
-3
5
2
-1
2
`
	expect := []bool{true, true, true, true, false, true}

	runSample(t, s, expect)
}
