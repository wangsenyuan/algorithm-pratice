package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	n := len(res)

	arr := readNNums(bufio.NewReader(strings.NewReader(expect)), n)

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 3
4 3
4 2
1 2
`
	expect := "0 0 0 0"

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2
3 4
6 4
2 3
1 3
3 5
`
	expect := "0 0 0 1 1 2"

	runSample(t, s, expect)
}
