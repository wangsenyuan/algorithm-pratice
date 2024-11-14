package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
2 0 2
3 2 3 3
4 7 0 1 5`
	expect := 16
	// 7 0 1 5 x = 3时，可以得到2， 然后用2，就可以得到3
	// 当x = 4时，怎么得到4呢？直接使用x，不参与任何操作
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
5 0 2 0 4 11
1 1
5 1 3 0 3 3`
	expect := 20
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 50
2 1 2
2 1 2`
	expect := 1281
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1
7 1 2 4 1 4 9 5`
	expect := 6
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4 114514
2 2 2
5 7 3 6 0 3
3 0 1 1
5 0 9 2 1 5`
	expect := 6556785365
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5 1919810
1 2
2 324003 0
3 1416324 2 1460728
4 1312631 2 0 1415195
5 1223554 192248 2 1492515 725556`
	expect := 1842836177961
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `3 0
3 0 1 2
3 0 2 3
3 0 2 3`
	expect := 4
	runSample(t, s, expect)
}
