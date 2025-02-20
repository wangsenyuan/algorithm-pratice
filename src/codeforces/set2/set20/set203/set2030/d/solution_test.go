package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
1 4 2 5 3
RLRLL
2
4
3`
	expect := []bool{true, true, false}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 5
1 5 2 4 8 3 6 7
RRLLRRRL
4
3
5
3
4`
	expect := []bool{false, true, false, false, false}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2
1 2 3 4 5 6
RLRLRL
4
5`
	expect := []bool{true, true}
	runSample(t, s, expect)
}
