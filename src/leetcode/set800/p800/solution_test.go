package p800

import "testing"

func TestSample1(t *testing.T) {
	color := "#09f166"
	expect := "#11ee66"
	res := similarRGB(color)
	if res != expect {
		t.Errorf("sample %s, expect %s, but got %s", color, expect, res)
	}
}

func TestSimliarity(t *testing.T) {
	res := similarity("09f166", "11ee66")
	if res != -73 {
		t.Errorf("simliarity, expect -73, but got %d", res)
	}
}

func TestSimliarity1(t *testing.T) {
	res := similarity("09f166", "227700")
	if res != -25913 {
		t.Errorf("simliarity, expect -73, but got %d", res)
	}
}

func TestStr(t *testing.T) {
	expect := "e"
	res := str(14)
	if expect != res {
		t.Errorf("str expect %s, but got %s", expect, res)
	}
}
