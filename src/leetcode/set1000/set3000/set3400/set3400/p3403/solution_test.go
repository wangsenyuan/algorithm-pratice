package p3403

import "testing"

func runSample(t *testing.T, word string, k int, expect string) {
	res := answerString(word, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "dbca"
	k := 2
	expect := "dbc"
	runSample(t, word, k, expect)
}

func TestSample2(t *testing.T) {
	word := "gggg"
	k := 4
	expect := "g"
	runSample(t, word, k, expect)
}

func TestSample3(t *testing.T) {
	word := "gh"
	k := 1
	expect := "gh"
	runSample(t, word, k, expect)
}

func TestSample4(t *testing.T) {
	word := "aann"
	k := 2
	expect := "nn"
	runSample(t, word, k, expect)
}
