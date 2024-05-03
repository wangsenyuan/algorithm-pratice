package main

import "testing"

func runSample(t *testing.T, record string, expect int) {
	res := solve(record)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	record := "sffs"
	expect := 1
	runSample(t, record, expect)
}

func TestSample2(t *testing.T) {
	record := "fsfs"
	expect := 2
	runSample(t, record, expect)
}
func TestSample3(t *testing.T) {
	record := "fsfsfssssfssffsffffsssfssffffffssssfsfsfsffffssffssssfsfsfsfsssfssfsffsssffffssfffffffsfsssffsssssffffsffsfssfffsssfssfsfffsffssfsffssssfsffssfffssfffsfsffs"
	expect := 666443222
	runSample(t, record, expect)
}
