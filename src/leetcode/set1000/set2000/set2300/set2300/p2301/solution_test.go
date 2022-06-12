package p2301

import "testing"

func runSample(t *testing.T, s string, sub string, mappings [][]byte, expect bool) {
	res := matchReplacement(s, sub, mappings)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "fool3e7bar"
	sub := "leet"
	mappings := [][]byte{
		[]byte("e3"),
		[]byte("t7"),
		[]byte("t8"),
	}
	expect := true
	runSample(t, s, sub, mappings, expect)
}

func TestSample2(t *testing.T) {
	s := "fooleetbar"
	sub := "f00l"
	mappings := [][]byte{
		[]byte("o0"),
	}
	expect := false
	runSample(t, s, sub, mappings, expect)
}

func TestSample3(t *testing.T) {
	s := "Fool33tbaR"
	sub := "leetd"
	mappings := [][]byte{
		[]byte("e3"),
		[]byte("t7"),
		[]byte("t8"),
		[]byte("db"),
		[]byte("pb"),
	}
	expect := true
	runSample(t, s, sub, mappings, expect)
}
