package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 4
1 4 2 3 5
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 4
4 2 4 2 4 2 4 2
`
	// 对于 i = 4,  sum(i) = 10, 12 % 4 - 4 = -4
	// 对于 i = 2,  sum(i) = 6, 6 % 4 - 2 = 0
	// (12 - 6) % 4 = 2 = 4 - 2
	// 但是 6 % 4 - 2 = 0
	//  12 % 4 - 4 = -4
	// 所以等于0的部分，不好处理
	// sum % k = i (0)
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 7
14 15 92 65 35 89 79 32 38 46
`
	expect := 8
	runSample(t, s, expect)
}
