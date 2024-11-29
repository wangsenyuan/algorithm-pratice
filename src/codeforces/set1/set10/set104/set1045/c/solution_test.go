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
	s := `5 7 2
1 2
1 3
1 4
2 3
2 4
3 4
1 5
1 4
2 5
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 11 4
1 2
2 3
3 4
4 5
1 3
1 6
3 5
3 7
4 7
5 7
6 8
1 5
2 4
6 7
3 8
`
	expect := []int{2, 2, 3, 3}
	runSample(t, s, expect)
}
