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

	m := len(res)
	reader = bufio.NewReader(strings.NewReader(expect))
	ans := readNNums(reader, m)

	if !reflect.DeepEqual(ans, res) {
		t.Fatalf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
1 2 3 4 5
1
2
3
4
5`
	expect := "0 1 1 1 2 "
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3
1 2 6 4 1 3
2
1
5`
	expect := "1 0 2"
	runSample(t, s, expect)
}
