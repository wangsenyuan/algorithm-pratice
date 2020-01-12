package p1320

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := minimumDistance(word)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", word, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "CAKE", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "HAPPY", 6)
}

func TestSample3(t *testing.T) {
	runSample(t, "NEW", 3)
}

func TestSample4(t *testing.T) {
	runSample(t, "YEAR", 7)
}

func TestSample5(t *testing.T) {
	runSample(t, "OPVUWZLCKTDPSUKGHAXIDWHLZFKNBDZEWHBSURTVCADUGTSDMCLDBTAGFWDPGXZBVARNTDICHCUJLNFBQOBTDWMGILXPSFWVGYBZVFFKQIDTOVFAPVNSQJULMVIERWAOXCKXBRI", 295)
}
