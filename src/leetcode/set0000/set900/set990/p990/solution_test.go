package p990

import "testing"

func runSample(t *testing.T, eqs []string, expect bool) {
	res := equationsPossible(eqs)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", eqs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	eqs := []string{"a==b", "b!=a"}
	expect := false
	runSample(t, eqs, expect)
}

func TestSample2(t *testing.T) {
	eqs := []string{"b==a", "a==b"}
	expect := true
	runSample(t, eqs, expect)
}

func TestSample3(t *testing.T) {
	eqs := []string{"a==b", "b==c", "a==c"}
	expect := true
	runSample(t, eqs, expect)
}

func TestSample4(t *testing.T) {
	eqs := []string{"a==b", "b!=c", "c==a"}
	expect := false
	runSample(t, eqs, expect)
}

func TestSample5(t *testing.T) {
	eqs := []string{"c==c", "b==d", "x!=z"}
	expect := true
	runSample(t, eqs, expect)
}

func TestSample6(t *testing.T) {
	eqs := []string{"b==b", "b==e", "e==c", "d!=e"}
	expect := true
	runSample(t, eqs, expect)
}

func TestSample7(t *testing.T) {
	eqs := []string{"a!=i", "g==k", "k==j", "k!=i", "c!=e", "a!=e", "k!=a", "a!=g", "g!=c"}
	expect := true
	runSample(t, eqs, expect)
}
