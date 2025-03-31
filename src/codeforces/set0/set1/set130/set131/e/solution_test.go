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
	s := `8 4
4 3
4 8
6 5
1 6
`
	expect := []int{0, 3, 0, 1, 0, 0, 0, 0, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 3
1 1
1 2
1 3
`
	expect := []int{0, 2, 1, 0, 0, 0, 0, 0, 0}
	runSample(t, s, expect)
}
