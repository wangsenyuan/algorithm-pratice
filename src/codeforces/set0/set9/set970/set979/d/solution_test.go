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
	s := `5
1 1
1 2
2 1 1 3
2 1 1 2
2 1 1 1
`
	expect := []int{2, 1, -1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
1 9
2 9 9 22
2 3 3 18
1 25
2 9 9 20
2 25 25 14
1 20
2 26 26 3
1 14
2 20 20 9
`
	expect := []int{9, 9, 9, -1, -1, -1}
	runSample(t, s, expect)
}
