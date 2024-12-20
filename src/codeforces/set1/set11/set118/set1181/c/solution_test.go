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
	s := `4 3
aaa
bbb
ccb
ddd
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
a
a
b
b
c
c
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10
aaaaarpppp
bbbbsssssu
cciiiiiiqq
ddmmgggggg
eeebbbbbbb
fffffqoooo
gxxxxrrrrr
hhfuuuuuuu
iiillqqqqq
jjjjjppwww
`
	expect := 138
	runSample(t, s, expect)
}
