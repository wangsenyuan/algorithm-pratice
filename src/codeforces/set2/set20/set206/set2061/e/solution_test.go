package main

import (
	"bufio"
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(bytes.NewReader([]byte(s)))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 3 2
7
5 6 3`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 2
5 6
5 6 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 2 5
3 1 4 1 5 9 2 6 5 3
7 8`
	expect := 11
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 1 0
1073741823 1073741823 1073741823 1073741823 1073741823
1073741823`
	expect := 5368709115
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 1 0
0
0`
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 5 28
811414488 12596772 787323389 393863925 260893684 178927067 89627776 47928964 1062489041 606234561
868114278 459010937 1024032511 869004158 532106167`
	expect := 872510130
	runSample(t, s, expect)
}
