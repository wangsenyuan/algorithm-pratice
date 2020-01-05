package p1309

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := freqAlphabets(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "10#11#12", "jkab")
}

func TestSample2(t *testing.T) {
	runSample(t, "1326#", "acz")
}

func TestSample3(t *testing.T) {
	runSample(t, "25#", "y")
}

func TestSample4(t *testing.T) {
	runSample(t, "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "abcdefghijklmnopqrstuvwxyz")
}
