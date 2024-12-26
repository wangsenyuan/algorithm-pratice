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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 5
1 4
4 5
5 6
6 7
3 5
`
	// 这个为啥是7呢？ 可以去掉2,3， 剩下的就是 1~5, 6~7
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
1 1
2 2
3 4
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
1 1
2 2
2 3
3 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 3
1 6
1 3
4 6
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 3
1 5
2 3
4 10
`
	expect := 7
	runSample(t, s, expect)
}
