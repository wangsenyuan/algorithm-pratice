package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1 2
1 2
Dmitry 1 1 1
Natalia 2 0
Natalia 1 1 1
Dmitry 2 0`
	expect := true

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4 3
1 3
1 2
2 3
3 4
Artem 1 1 4
Dmitry 1 1 2
Edvard 4 2 1 3
Artem 2 0
Dmitry 1 0
Edvard 4 4 1 2 3 4
`
	expect := false

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 3
1 2
2 3
3 1
a 1 1 1
b 2 1 3
c 3 1 2
b 1 1 2
c 2 1 3
a 3 1 1`
	expect := false

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 1 1
2 1
absgdf 2 1 1
absgdf 1 1 1`
	expect := true

	runSample(t, s, expect)
}
