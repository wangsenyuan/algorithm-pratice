package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	r := bufio.NewReader(strings.NewReader(s))

	res := process(r)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
7
1 2 3 4 5 6 7
3
1 3 2
3
4 5 3
15
9 3 8 1 6 5 3 8 2 1 4 2 9 4 7
`
	expect := []int{3, 1, 0, 8}
	runSample(t, s, expect)
}
