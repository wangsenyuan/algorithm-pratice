package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	reader := bufio.NewReader(strings.NewReader(expect))
	ans := readNNums(reader, len(res))

	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
2 3
3 4
3 5
3
5 1
3 1
2 0`
	expect := "2 1 2"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9
8 1
1 7
1 4
7 3
4 9
3 2
1 5
3 6
7
6 0
2 3
6 2
8 2
2 4
9 2
6 3`
	expect := "0 5 2 4 5 5 5"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
2 1
2 5
2 4
5 6
4 3
3
3 1
1 3
6 5`
	expect := "1 3 4"
	runSample(t, s, expect)
}