package p2060

import "testing"

func runSample(t *testing.T, s1, s2 string, expect bool) {
	res := possiblyEquals(s1, s2)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "internationalization"
	s2 := "i18n"
	expect := true

	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "l123e"
	s2 := "44"
	expect := true

	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "a5b"
	s2 := "c5b"
	expect := false

	runSample(t, s1, s2, expect)
}

func TestSample4(t *testing.T) {
	s1 := "112s"
	s2 := "g841"
	expect := true

	runSample(t, s1, s2, expect)
}
