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
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
Kuroo
Shiro
Katie
`
	expect := users[0]
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
treasurehunt
threefriends
hiCodeforces
`
	expect := users[1]
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
aaaaaaaaaa
AAAAAAcAAA
bbbbbbzzbb
`
	expect := users[1]
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `61
bzqiqprzfwddqwctcrhnkqcsnbmcmfmrgaljwieajfouvuiunmfbrehxchupmsdpwilwu
jyxxujvxkwilikqeegzxlyiugflxqqbwbujzedqnlzucdnuipacatdhcozuvgktwvirhs
tqiahohijwfcetyyjlkfhfvkhdgllxmhyyhhtlhltcdspusyhwpwqzyagtsbaswaobwub
`
	expect := "Katie"
	runSample(t, s, expect)
}
