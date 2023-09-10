package p2848

import "testing"

func runSample(t *testing.T, s string, x string, k int64, expect int) {
	res := numberOfWays(s, x, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	x := "cdab"
	var k int64 = 2
	expect := 2
	runSample(t, s, x, k, expect)
}

func TestSample2(t *testing.T) {
	s := "ababab"
	x := "ababab"
	var k int64 = 1
	expect := 2
	runSample(t, s, x, k, expect)
}

func TestSample3(t *testing.T) {
	s := "ttttttt"
	x := "ttttttt"
	var k int64 = 5
	expect := 7776
	runSample(t, s, x, k, expect)
}
