package p2070

import "testing"

func runSample(t *testing.T, encoded string, rows int, expect string) {
	res := decodeCiphertext(encoded, rows)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	encoded := "ch   ie   pr"
	rows := 3
	expect := "cipher"
	runSample(t, encoded, rows, expect)
}

func TestSample2(t *testing.T) {
	encoded := "iveo    eed   l te   olc"
	rows := 4
	expect := "i love leetcode"
	runSample(t, encoded, rows, expect)
}

func TestSample3(t *testing.T) {
	encoded := "coding"
	rows := 1
	expect := "coding"
	runSample(t, encoded, rows, expect)
}

func TestSample4(t *testing.T) {
	encoded := " b  ac"
	rows := 2
	expect := " abc"
	runSample(t, encoded, rows, expect)
}
