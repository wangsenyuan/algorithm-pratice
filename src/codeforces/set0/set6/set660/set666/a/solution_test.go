package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abacabaca"
	expect := []string{"aca", "ba", "ca"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abaca"
	runSample(t, s, nil)
}

func TestSample3(t *testing.T) {
	s := "affviytdmexpwfqplpyrlniprbdphrcwlboacoqec"
	expect := []string{
		"ac", "aco", "bd", "bdp", "bo", "boa", "co", "coq", "cw", "cwl", "dm", "dme", "dp", "dph", "ec", "ex", "exp", "fq", "fqp", "hr", "hrc", "ip", "ipr", "lb", "lbo", "ln", "lni", "lp", "lpy", "me", "mex", "ni", "nip", "oa", "oac", "oq", "ph", "phr", "pl", "plp", "pr", "prb", "pw", "pwf", "py", "pyr", "qec", "qp", "qpl", "rb", "rbd", "rc", "rcw", "rl", "rln", "td", "tdm", "wf", "wfq", "wl", "wlb", "xp", "xpw", "yr", "yrl", "yt", "ytd",
	}
	runSample(t, s, expect)
}
