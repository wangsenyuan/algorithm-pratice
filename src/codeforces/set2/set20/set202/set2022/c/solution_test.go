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
	s := `3
AAA
AJJ`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
JAJAJJ
JJAJAJ`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
AJJJAJ
AJJAAA`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
AJJJJAJAJ
JAAJJJJJA`
	expect := 2
	runSample(t, s, expect)
}