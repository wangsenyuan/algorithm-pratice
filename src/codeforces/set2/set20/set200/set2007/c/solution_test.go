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
	s := `4 5 5
1 3 4 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2 3
1 3 4 6`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 7 7
1 1 2 6`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 15 9
1 9 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 18 12
1 4 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `7 27 36
33 13 23 12 35 24 41`
	expect := 5
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `10 6 9
15 5 6 9 8 2 12 15 3 8`
	expect := 1
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `6 336718728 709848696
552806726 474775724 15129785 371139304 178408298 13106071`
	expect := 17
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `6 335734893 671469786
138885253 70095920 456876775 9345665 214704906 375508929`
	expect := 205359241
	runSample(t, s, expect)
}
