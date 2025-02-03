package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	// n := len(res)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	ans := buf.String()
	if ans != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
abcde
edcba
1 5
1 4
3 3`
	expect := `0
1
0
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
zzde
azbe
1 3
1 4`
	expect := `2
2
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3
uwuwuw
wuwuwu
2 4
1 3
1 6`
	expect := `1
1
0
`
	runSample(t, s, expect)
}
