package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5
5 4 3 2 1
`
	expect := []int{0, 1, 3, 6, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 3
`
	expect := []int{0, 0, 0}
	runSample(t, s, expect)
}
