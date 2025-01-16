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
	s := `5 2
1 2 15 15 15
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
36 4 1 25 9 16
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `50 2
4068 1859 72 3173 3163 1640 437 832 2262 1193 3693 3930 3626 3795 1296 377 566 3888 606 2561 2225 3812 1936 1280 3982 345 3485 4930 1303 1652 407 1342 359 2669 668 3382 1463 259 3813 2915 58 2366 2930 2904 1536 3931 91 194 3447 4892
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50 1
849 1987 4384 4180 1393 97 1596 2522 451 4037 2186 2421 633 4882 4044 225 1569 4998 4915 2296 3509 1788 4713 4774 810 4732 2031 907 514 3342 3296 1973 4858 315 3140 3651 3212 4996 1390 1323 862 872 438 1533 1378 1411 3232 1565 3724 3753
`
	expect := 2
	runSample(t, s, expect)
}
