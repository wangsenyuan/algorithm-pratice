package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	ans := process(bufio.NewReader(strings.NewReader(s)))

	reader := bufio.NewReader(strings.NewReader(expect))

	res := readNNums(reader, len(ans))

	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
1 3
2 4
2 5
2
1 4
3 5
`
	expect := "2 1 1 1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 4
4 5
1 4
2 4
3
2 3
1 3
3 5
`
	expect := "3 1 1 1"
	runSample(t, s, expect)
}
