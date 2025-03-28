package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
Ann
Anna
Sabrina
John
Petrov
Ivanova
Stoltz
Abacaba
`
	expect := "Ann Abacaba, Anna Ivanova, John Petrov, Sabrina Stoltz"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
Aa
Ab
Ac
Ba
Ad
Ae
Bb
Bc
`
	expect := "Aa Ad, Ab Ae, Ac Bb, Ba Bc"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
Ccbcbaba
Ab
Cbcbcc
Cac
Cba
Cacabbabbb
Aacacbcac
Aacbc
Bbccc
Bbbb
`
	expect := "Ab Aacacbcac, Cac Aacbc, Cba Bbbb, Cbcbcc Bbccc, Ccbcbaba Cacabbabbb"
	runSample(t, s, expect)
}
