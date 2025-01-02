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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `9 6
1 1 1 3 5 3 5 7
3 1
1 5
3 4
7 3
1 8
1 9
`
	expect := []int{3, 6, 8, -1, 9, 4}
	runSample(t, s, expect)
}
