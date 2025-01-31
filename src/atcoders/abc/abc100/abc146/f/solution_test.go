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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `9 3
0001000100
`
	expect := []int{1, 3, 2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
011110
`

	runSample(t, s, nil)
}

func TestSample3(t *testing.T) {
	s := `6 6
0101010
`
	expect := []int{6}
	runSample(t, s, expect)
}
