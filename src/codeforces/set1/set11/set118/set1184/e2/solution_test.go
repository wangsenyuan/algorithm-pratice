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
	s := `3 3
1 2 8
2 3 3
3 1 4
`
	expect := []int{4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2 8
2 3 3
3 1 4
`
	expect := []int{4}
	runSample(t, s, expect)
}
