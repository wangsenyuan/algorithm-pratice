package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	ans := readNNums(bufio.NewReader(strings.NewReader(expect)), len(res))

	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10
1 2 3 4 5 6 7 8 9 10`
	expect := "1 3 8 13 46 59 126 149 1174 1311"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11
1 6 9 4 7 4 4 10 3 2 3`
	expect := "1 7 22 26 70 74 150 1303 1306 1308 1568"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
527792568 502211460 850237282 374773208`
	expect := "527792568 83665723 399119771 773892979"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
10 5 7 7 4 10 10 4 2`
	expect := "10 15 24 31 35 105 190 194 292 "
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `16
2 1 10 1 5 5 1 9 2 4 4 4 1 8 3 5`
	expect := "2 3 22 23 28 33 34 55 57 64 158 543 544 4129 12322 20517"
	runSample(t, s, expect)
}
