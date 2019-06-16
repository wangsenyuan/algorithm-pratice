package p726

import "testing"

func TestSample1(t *testing.T) {
	str := "H2O"
	res := countOfAtoms(str)
	if res != str {
		t.Errorf("atoms of %s should be %s, but is %s", str, str, res)
	}
}

func TestSample2(t *testing.T) {
	str := "Mg(OH)2"
	res := countOfAtoms(str)
	expect := "H2MgO2"
	if res != expect {
		t.Errorf("atoms of %s should be %s, but is %s", str, expect, res)
	}
}

func TestSample3(t *testing.T) {
	str := "K4(ON(SO3)2)2"
	res := countOfAtoms(str)
	expect := "K4N2O14S4"
	if res != expect {
		t.Errorf("atoms of %s should be %s, but is %s", str, expect, res)
	}
}
