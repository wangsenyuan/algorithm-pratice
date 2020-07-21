package p772

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := calculate(s)
	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1+1", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, " 6-4 / 2 ", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "2*(5+5*2)/3+(6/2+8)", 21)
}

func TestSample4(t *testing.T) {
	runSample(t, "(2+6* 3+5- (3*14/7+2)*5)+3", -12)
}
