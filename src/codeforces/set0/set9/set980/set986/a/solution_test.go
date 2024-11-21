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
	s := `5 5 4 3
1 2 4 3 2
1 2
2 3
3 4
4 1
4 5
`
	expect := []int{2, 2, 2, 2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 6 3 2
1 2 3 3 2 2 1
1 2
2 3
3 4
2 5
5 6
6 7
`
	expect := []int{1, 1, 1, 2, 2, 1, 1}
	runSample(t, s, expect)
}
