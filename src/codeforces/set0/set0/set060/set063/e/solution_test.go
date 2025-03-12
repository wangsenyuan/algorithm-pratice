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
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `  . . .
 . . O .
. . O O .
 . . . .
  . . .
`
	expect := "Lillebror"

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `  . . .
 . . . O
. . . O .
 O . O .
  . O .
`
	expect := "Karlsson"

	runSample(t, s, expect)
}
