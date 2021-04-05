package a

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := truncateSentence(s, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "Hello how are you Contestant"
	k := 4
	expect := "Hello how are you"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "What is the solution to this problem"
	k := 4
	expect := "What is the solution"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "chopper is not a tanuki"
	k := 5
	expect := "chopper is not a tanuki"
	runSample(t, s, k, expect)
}
