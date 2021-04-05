package p1813

import "testing"

func runSample(t *testing.T, s1, s2 string, expect bool) {
	res := areSentencesSimilar(s1, s2)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "eTUny i b R UFKQJ EZx JBJ Q xXz"
	s2 := "eTUny i R EZx JBJ xXz"
	expect := false
	runSample(t, s1, s2, expect)
}

func TestSample2(t *testing.T) {
	s1 := "My name is Haley"
	s2 := "My Haley"
	expect := true
	runSample(t, s1, s2, expect)
}

func TestSample3(t *testing.T) {
	s1 := "Lucccy"
	s2 := "Luccy"
	expect := false
	runSample(t, s1, s2, expect)
}

func TestSample4(t *testing.T) {
	s1 := "She is Luccy"
	s2 := "Luccy"
	expect := true
	runSample(t, s1, s2, expect)
}
