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
	s := `2
2 3
5
2 1 1
2 2 1
1 1 1
2 1 2
2 2 0
`
	expect := []int{4, 3, 3, 2, 3}
	runSample(t, s, expect)
}
