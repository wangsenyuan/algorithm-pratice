package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	ans := process(reader)

	var buf bytes.Buffer
	for _, x := range ans {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}

	s = buf.String()

	if s != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, s)
	}
}

func TestSample1(t *testing.T) {
	s := `5 4
1 2 3 2 4
2 3 1 4 2
1 3 1 3
1 2 3 5
1 4 2 5
1 5 1 5
`
	expect := `Yes
No
No
Yes
`
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `4 4
4 4 4 4
4 4 4 4
1 2 2 3
3 3 1 1
1 3 1 4
1 4 2 3
`
	expect := `Yes
Yes
No
No
`
	runSample(t, s, expect)
}
