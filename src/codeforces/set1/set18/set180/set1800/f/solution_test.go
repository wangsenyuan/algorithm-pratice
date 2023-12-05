package main

import "testing"

func runSample(t *testing.T, s []string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"ftl",
		"abcdefghijklmnopqrstuvwxy",
		"abcdeffghijkllmnopqrsttuvwxy",
		"ffftl",
		"aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyy",
		"thedevid",
		"bcdefghhiiiijklmnopqrsuwxyz",
		"gorillasilverback",
		"abcdefg",
		"ijklmnopqrstuvwxyz",
	}
	expect := 5
	runSample(t, s, expect)
}
