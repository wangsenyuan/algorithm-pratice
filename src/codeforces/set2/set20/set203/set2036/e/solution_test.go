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
	s := `3 4 4
1 3 5 9
4 6 5 3
2 1 2 7
3
1 > 4
2 < 8
1 < 6
2
1 < 8
2 > 8
1
3 > 5
2
4 > 8
1 < 8
`
	expect := []int{2, -1, 3, 1}
	runSample(t, s, expect)
}
