package p2272

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := largestVariance(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aababbb"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaaa"
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "hhsxhhorrs"
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "bbc"
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "ykudzhiixwttnvtesiwnbcjmsydidttiyabbwzlfbmmycwjgzwhbtvtxyvkkjgfehaypiygpstkhakfasiloaveqzcywsiujvixcdnxpvvtobxgroznswwwipypwmdhldsoswrzyqthaqlbwragjrqwjxgmftjxqugoonxadazeoxalmccfeyqtmoxwbnphxih"
	expect := 12
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "abb"
	expect := 1
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "aabaaa"
	expect := 4
	runSample(t, s, expect)
}