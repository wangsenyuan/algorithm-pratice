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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `9
2 1
4 5
9
2 1
3 6
9
4 2
1 2 2 2
3 4
4 2
1 1 3 3
3 5
4 2
1 2 3 4
3 5
5 5
1 2 3 4 5
5 4 3 2 1
4 2
1 1 1 1
1 1
4 4
1 1 1 1
1 1 1 2
1 1
1
1000000000
`
	expect := []bool{true, false, true, true, false, true, false, false, false}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
18 8
1 2 10 3 3 2 3 3 2 1 1 1 5 2 1 1 2 2
7 9 1 6 5 2 10 5`
	expect := []bool{true}
	runSample(t, s, expect)
}
