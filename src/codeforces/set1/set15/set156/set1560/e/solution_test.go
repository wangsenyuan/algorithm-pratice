package main

import "testing"

type answer struct {
	ok  bool
	s   string
	ord string
}

func runSample(t *testing.T, x string, expect answer) {
	ok, pref, ord := solve(x)

	res := answer{ok, pref, ord}

	if expect != res {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abacabaaacaac"
	expect := answer{
		ok:  true,
		s:   "abacaba",
		ord: "bac",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "nowyouknowthat"
	expect := answer{
		ok:  false,
		s:   "",
		ord: "",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "polycarppoycarppoyarppyarppyrpprppp"
	expect := answer{
		ok:  true,
		s:   "polycarp",
		ord: "lcoayrp",
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "isi"
	expect := answer{
		ok:  true,
		s:   "is",
		ord: "si",
	}
	runSample(t, s, expect)
}
func TestSample5(t *testing.T) {
	s := "everywherevrywhrvryhrvrhrvhv"
	expect := answer{
		ok:  true,
		s:   "everywhere",
		ord: "ewyrhv",
	}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "haaha"
	expect := answer{
		ok:  false,
		s:   "",
		ord: "",
	}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	// 字符串的顺序也有关系
	s := "qweqeewew"
	expect := answer{
		ok:  false,
		s:   "",
		ord: "",
	}
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	// 字符串的顺序也有关系
	s := "v"
	expect := answer{
		ok:  true,
		s:   "v",
		ord: "v",
	}
	runSample(t, s, expect)
}
