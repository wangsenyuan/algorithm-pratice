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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ab
bc
cd
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abacaba
abaaba
x
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abc
abc
ab
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `xufuzdlsjxmevrtessfbwlnzzclcqwevnnucxyvhngnxhcbdfwq
wlwobhnmmgtfolfaeckbrnnglylydxtgtvrlmeeszoiuatzzzxufuzdlsjxmevrt
brnnglylydxtgtvrlmeeszoiuatzzzx
`
	expect := 100
	runSample(t, s, expect)
}
