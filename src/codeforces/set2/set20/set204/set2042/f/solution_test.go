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
	s := `7
3 -1 4 -3 2 4 0
0 6 1 0 -3 -2 -1
6
3 1 7
1 2 0
3 3 6
2 5 -3
1 3 2
3 1 5
`
	expect := []int{18, 7, 16}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
2 -1 -3 -2 0 4 5 6 2 5
2 -4 -5 -1 6 2 5 -6 4 2
10
3 6 7
1 10 -2
3 5 7
3 2 8
2 1 -5
2 7 4
3 1 3
3 3 8
3 2 3
1 4 4
`
	expect := []int{23, 28, 28, -17, 27, -22}
	runSample(t, s, expect)
}
