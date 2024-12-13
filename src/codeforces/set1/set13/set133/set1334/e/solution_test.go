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
	s := `12
3
4 4
12 1
3 4
`
	expect := []int{1, 3, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `288807105787200
4
46 482955026400
12556830686400 897
414 12556830686400
4443186242880 325
`
	expect := []int{547558588, 277147129, 457421435, 702277623}
	runSample(t, s, expect)
}
